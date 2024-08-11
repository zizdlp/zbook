package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func CreateRandomVerification(t *testing.T) Verification {
	user := createRandomUser(t)

	verificationType := util.RandomVerificationType()

	arg := CreateVerificationParams{VerificationUrl: util.RandomString(32), UserID: user.UserID, VerificationType: verificationType}
	verification, err := testStore.CreateVerification(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, verification.IsUsed, false)
	require.Equal(t, verification.UserID, user.UserID)
	require.Equal(t, verification.VerificationType, verificationType)
	return verification
}
func TestCreateVerification(t *testing.T) {
	CreateRandomVerification(t)
}
func TestGetVerification(t *testing.T) {
	verification := CreateRandomVerification(t)
	v2, err := testStore.GetVerification(context.Background(), verification.VerificationUrl)
	require.NoError(t, err)
	require.Equal(t, v2.UserID, verification.UserID)
	require.Equal(t, v2.IsUsed, verification.IsUsed)
	require.Equal(t, v2.CreatedAt, verification.CreatedAt)
	require.Equal(t, v2.VerificationType, verification.VerificationType)
}
func TestMarkVerificationAsUsed(t *testing.T) {
	verification := CreateRandomVerification(t)
	v2, err := testStore.MarkVerificationAsUsed(context.Background(), verification.VerificationUrl)
	require.NoError(t, err)
	require.Equal(t, v2.UserID, verification.UserID)
	require.Equal(t, v2.IsUsed, true)
	require.Equal(t, v2.CreatedAt, verification.CreatedAt)
	require.Equal(t, v2.VerificationType, verification.VerificationType)
}
