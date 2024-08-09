package db

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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

func TestGetListSessionCount(t *testing.T) {
	user := createRandomUser(t)
	createRandomSession(t, user)
	count1, err := testStore.GetListSessionCount(context.Background())
	require.NoError(t, err)
	createRandomSession(t, user)
	count2, err := testStore.GetListSessionCount(context.Background())
	require.NoError(t, err)
	require.Equal(t, count2, count1+1)
}
func TestListSession(t *testing.T) {
	user := createRandomUser(t)
	session1 := createRandomSession(t, user)
	session2 := createRandomSession(t, user)
	arg := ListSessionParams{
		Limit:  2,
		Offset: 0,
	}
	sessions, err := testStore.ListSession(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, 2, len(sessions))
	require.Equal(t, sessions[0].SessionID, session2.SessionID)
	require.Equal(t, sessions[1].SessionID, session1.SessionID)
}

func TestGetQuerySessionCount(t *testing.T) {
	user := createRandomUser(t)
	createRandomSession(t, user)
	count1, err := testStore.GetQuerySessionCount(context.Background(), user.Username)
	require.NoError(t, err)
	createRandomSession(t, user)
	count2, err := testStore.GetQuerySessionCount(context.Background(), user.Username)
	require.NoError(t, err)
	require.Equal(t, count2, count1+1)
}
func TestQuerySession(t *testing.T) {
	user := createRandomUser(t)
	session1 := createRandomSession(t, user)
	session2 := createRandomSession(t, user)
	arg := QuerySessionParams{
		Limit:  2,
		Offset: 0,
		Query:  user.Username,
	}
	sessions, err := testStore.QuerySession(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, 2, len(sessions))
	require.Equal(t, sessions[0].SessionID, session2.SessionID)
	require.Equal(t, sessions[1].SessionID, session1.SessionID)
}

func TestGetDailyActiveUserCount(t *testing.T) {
	user := createRandomUser(t)
	createRandomSession(t, user)
	timezone := "America/New_York"
	arg := GetDailyActiveUserCountParams{
		Timezone:     timezone,
		IntervalDays: pgtype.Text{String: "7", Valid: true},
	}
	count1, err := testStore.GetDailyActiveUserCount(context.Background(), arg)
	require.NoError(t, err)
	require.True(t, len(count1) > 0)

	user2 := createRandomUser(t)
	createRandomSession(t, user2)
	count2, err := testStore.GetDailyActiveUserCount(context.Background(), arg)
	require.NoError(t, err)
	require.True(t, len(count2) > 0)
	require.Equal(t, count2[0].ActiveUsersCount, count1[0].ActiveUsersCount+1)
}
