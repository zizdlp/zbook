package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomUsername(),
		Email:          util.RandomEmail(),
		HashedPassword: hashedPassword,
	}
	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.CreatedAt)
	require.Equal(t, user.UserRole, util.UserRole)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestUpdateUserBasicInfo(t *testing.T) {
	user1 := createRandomUser(t)

	motto := util.RandomString(6)
	hashedPassword, _ := util.HashPassword(util.RandomString(6))

	arg := UpdateUserBasicInfoParams{
		Username: user1.Username,
		Motto: pgtype.Text{
			String: motto,
			Valid:  true,
		},
		HashedPassword: pgtype.Text{
			String: hashedPassword,
			Valid:  true,
		},
		UserRole: pgtype.Text{
			String: util.RandomUserRole(),
			Valid:  util.RandomBool(),
		},
		Onboarding: util.RandomPGBool(),
		Blocked:    util.RandomPGBool(),
		Verified:   util.RandomPGBool(),
	}
	user2, err := testStore.UpdateUserBasicInfo(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user2.Username, arg.Username)
	require.Equal(t, user2.HashedPassword, hashedPassword)
	require.Equal(t, user2.Motto, motto)
}
func TestUpdateUnreadCount(t *testing.T) {
	user1 := createRandomUser(t)
	err := testStore.UpdateUnreadCount(context.Background(), user1.UserID)
	require.NoError(t, err)
}
func TestResetnreadCount(t *testing.T) {
	user1 := createRandomUser(t)
	err := testStore.ResetUnreadCount(context.Background(), user1.Username)
	require.NoError(t, err)
}
