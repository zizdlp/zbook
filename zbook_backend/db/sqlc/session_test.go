package db

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func createRandomSession(t *testing.T, user User) Session {
	arg := CreateSessionParams{
		SessionID:    uuid.New(),
		UserID:       user.UserID,
		RefreshToken: util.RandomString(32),
		UserAgent:    util.RandomString(6),
		ClientIp:     util.RandomString(6),
		ExpiresAt:    time.Now(),
	}

	session, err := testStore.CreateSession(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, session)
	require.Equal(t, arg.SessionID, session.SessionID)
	require.Equal(t, arg.UserAgent, session.UserAgent)
	require.Equal(t, arg.UserID, session.UserID)
	require.Equal(t, arg.RefreshToken, session.RefreshToken)
	require.Equal(t, arg.ClientIp, session.ClientIp)
	require.NotZero(t, user.CreatedAt)
	return session
}

func TestCreateSession(t *testing.T) {
	user := createRandomUser(t)
	createRandomSession(t, user)
}

func TestGetSession(t *testing.T) {
	user := createRandomUser(t)
	session := createRandomSession(t, user)

	session2, err := testStore.GetSession(context.Background(), session.SessionID)
	require.NoError(t, err)
	require.NotEmpty(t, session2)
	require.Equal(t, session2.SessionID, session.SessionID)
}
