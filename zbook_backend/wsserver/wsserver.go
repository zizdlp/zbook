package wsserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/zizdlp/zbook/token"
	"github.com/zizdlp/zbook/util"
	"golang.org/x/sync/errgroup"
)

type message struct {
	Username           string `json:"username"`
	UnreadMessageCount int    `json:"unread_count"`
}

var (
	websocketMap sync.Map // map[string][]*websocket.Conn
)

func WebSocketServer(ctx context.Context, waitGroup *errgroup.Group, config util.Config) {
	// Start a goroutine to handle incoming notifications.
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Error().Err(err).Msg("failed to create token maker")
		return
	}

	httpServer := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/ws":
				handleWebSocket(w, r, tokenMaker)
			case "/ws/connections":
				ListWebSocketConnections(w, r, tokenMaker)
			default:
				http.NotFound(w, r)
			}
		}),
		Addr: config.WEBSOCKETServerAddress,
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("start WebSocket server at %s", httpServer.Addr)
		err = httpServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			log.Error().Err(err).Msg("WebSocket server failed to serve")
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown WebSocket server")

		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("failed to shutdown WebSocket server")
			return err
		}

		log.Info().Msg("WebSocket server is stopped")
		return nil
	})
}

func ListenWebSocket(pool *pgxpool.Pool) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("Error acquiring connection")
		os.Exit(1)
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), "LISTEN unread_count_change")
	if err != nil {
		log.Error().Err(err).Msg("Error listening to chat channel")
		os.Exit(1)
	}

	for {
		notification, err := conn.Conn().WaitForNotification(context.Background())
		if err != nil {
			if err.Error() == "unexpected EOF" {
				// Handle the error by reconnecting to the server.
				log.Warn().Msg("Reconnecting to the websocket server...")
				conn.Release()
				conn, err = pool.Acquire(context.Background())
				if err != nil {
					log.Error().Err(err).Msg("Error acquiring connection")
					os.Exit(1)
				}
				_, err = conn.Exec(context.Background(), "LISTEN unread_count_change")
				if err != nil {
					log.Error().Err(err).Msg("Error listening to chat channel")
					os.Exit(1)
				}
				continue
			} else {
				log.Error().Msgf("Error waiting for notification: %s", err)
				os.Exit(1)
			}
		}
		var msg message
		err = json.Unmarshal([]byte(notification.Payload), &msg)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal json payload")
			continue
		}
		// Send the notification to the appropriate WebSocket connection.
		sendWebSocketMessage(msg.Username, msg.UnreadMessageCount)
	}
}

func sendWebSocketMessage(username string, unreadMessageCount int) {
	// Get the WebSocket connections for the specified user.
	conns, ok := websocketMap.Load(username)
	if !ok {
		return
	}

	// Construct the message payload.
	payload := fmt.Sprintf(`{"username":"%s","unread_count":%d}`, username, unreadMessageCount)

	// Send the message over all WebSocket connections.
	for _, conn := range conns.([]*websocket.Conn) {
		err := conn.WriteMessage(websocket.TextMessage, []byte(payload))
		if err != nil {
			log.Error().Err(err).Msgf("failed to write websocket message for user: %s", username)
			// If there is an error, close the connection and remove it from the map
			conn.Close()
			removeWebSocketConnection(username, conn)
		}
	}
}
func removeWebSocketConnection(username string, conn *websocket.Conn) {
	// Retrieve the existing connections for the user.
	conns, ok := websocketMap.Load(username)
	if !ok {
		return
	}

	// Create a new slice excluding the specified connection.
	var updatedConns []*websocket.Conn
	for _, existingConn := range conns.([]*websocket.Conn) {
		if existingConn != conn {
			updatedConns = append(updatedConns, existingConn)
		}
	}

	// Update the map or delete the entry if no connections remain.
	if len(updatedConns) > 0 {
		websocketMap.Store(username, updatedConns)
	} else {
		websocketMap.Delete(username)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request, tokenMaker token.Maker) {
	// Get the user ID from the request.
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}

	// Upgrade the HTTP connection to a WebSocket connection.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to upgrade http to WebSocket")
		return
	}

	_, accessToken, err := conn.ReadMessage()
	if err != nil {
		log.Error().Err(err).Msg("failed to read access token message")
		conn.Close()
		return
	}

	payload, err := tokenMaker.VerifyToken(string(accessToken))
	if err != nil {
		log.Error().Err(err).Msgf("websocket:failed to verify token for user:%s", username)
		conn.Close()
		return
	}

	if payload.Username != username {
		log.Error().Err(err).Msg("websocket:permission denied account not match")
		conn.Close()
		return
	}

	// Retrieve the existing connections for the user, or create a new slice if none exist.
	var conns []*websocket.Conn
	if existingConns, ok := websocketMap.Load(username); ok {
		conns = existingConns.([]*websocket.Conn)
	}
	conns = append(conns, conn)

	// Store the updated slice of connections.
	websocketMap.Store(username, conns)

	log.Info().Msgf("websocket:connect to user: %s", username)

	// Start a new goroutine for each WebSocket connection.
	go func() {
		defer conn.Close()

		// Loop to handle incoming WebSocket messages.
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				// Log the reason for the connection closure
				if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
					log.Info().Msgf("WebSocket connection closed for user: %s, reason: %v", username, err)
				} else {
					log.Warn().Msgf("WebSocket connection error for user: %s, error: %v", username, err)
				}
				removeWebSocketConnection(username, conn)
				break
			}
		}
	}()
}

func ListWebSocketConnections(w http.ResponseWriter, r *http.Request, tokenMaker token.Maker) {
	connections := make(map[string]int) // count of connections per user
	accessToken := r.URL.Query().Get("access_token")
	if accessToken == "" {
		http.Error(w, "Invalid access_token", http.StatusBadRequest)
		return
	}

	// Verify the access token.
	payload, err := tokenMaker.VerifyToken(accessToken)
	if err != nil {
		log.Error().Err(err).Msg("websocket: failed to verify token")
		http.Error(w, "Invalid access_token", http.StatusUnauthorized)
		return
	}

	// Check if the user has admin privileges.
	if payload.Role != util.AdminRole {
		log.Error().Msg("websocket: permission denied, only admin account can use this API")
		http.Error(w, "Permission denied", http.StatusForbidden)
		return
	}

	// Iterate over the websocketMap to collect all active connections.
	websocketMap.Range(func(key, value interface{}) bool {
		username, ok := key.(string)
		if ok {
			connections[username] = len(value.([]*websocket.Conn))
		}
		return true
	})

	// Convert the map to JSON.
	response, err := json.Marshal(connections)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type and write the response.
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
