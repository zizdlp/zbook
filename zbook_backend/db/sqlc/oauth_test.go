package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func createRandomOAuth(t *testing.T) Oauth {
	randomOAuthType := util.RandomOAuth()
	user := createRandomUser(t)
	arg := CreateOAuthParams{
		UserID:    user.UserID,
		OauthType: randomOAuthType,
		AppID:     util.RandomString(32),
	}
	oauth, err := testStore.CreateOAuth(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, oauth.UserID, user.UserID)
	require.Equal(t, oauth.OauthType, randomOAuthType)
	return oauth
}
func createOAuthUserOAuth(t *testing.T, user User, oauthType string) Oauth {

	arg := CreateOAuthParams{
		UserID:    user.UserID,
		OauthType: oauthType,
		AppID:     util.RandomString(32),
	}
	oauth, err := testStore.CreateOAuth(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, oauth.UserID, user.UserID)
	require.Equal(t, oauth.OauthType, oauthType)
	return oauth
}

func TestCreateOAuth(t *testing.T) {
	createRandomOAuth(t)
}

func TestGetOAuthUser(t *testing.T) {
	oauth := createRandomOAuth(t)
	arg := GetOAuthUserParams{
		OauthType: oauth.OauthType,
		AppID:     oauth.AppID,
	}
	oauthRow, err := testStore.GetOAuthUser(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, oauthRow.UserID, oauth.UserID)
	require.Equal(t, oauthRow.OauthType, oauth.OauthType)
	require.Equal(t, oauthRow.AppID, oauth.AppID)
}
func TestCheckOAuthStatus(t *testing.T) {
	user := createRandomUser(t)
	createOAuthUserOAuth(t, user, util.OAuthTypeGithub)
	createOAuthUserOAuth(t, user, util.OAuthTypeGoogle)
	status, err := testStore.CheckOAuthStatus(context.Background(), user.UserID)
	require.NoError(t, err)
	require.Equal(t, status.GithubStatus, true)
	require.Equal(t, status.GoogleStatus, true)
}

func TestDeleteOAuth(t *testing.T) {
	oauth := createRandomOAuth(t)
	arg := DeleteOAuthParams{
		UserID:    oauth.UserID,
		OauthType: oauth.OauthType,
	}
	_, err := testStore.DeleteOAuth(context.Background(), arg)
	require.NoError(t, err)

	arg_get := GetOAuthUserParams{
		OauthType: oauth.OauthType,
		AppID:     oauth.AppID,
	}
	oauthRow, err := testStore.GetOAuthUser(context.Background(), arg_get)
	require.Error(t, err)
	require.Empty(t, oauthRow)
}
