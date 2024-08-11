package storage

import (
	"context"
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func TestMinio(t *testing.T) {
	if testing.Short() {
		fmt.Println("***** TestMinio is ignored *****")
		t.Skip()
	}
	minioClient, err := GetMinioClient()
	require.NoError(t, err)
	require.NoError(t, err)
	bucketName := util.RandomString(32)
	err = CreateBucket(minioClient, context.Background(), bucketName)
	require.NoError(t, err)
	username := util.RandomUsername()
	avatarData, err := GenerateRandomBytes(32)
	require.NoError(t, err)
	err = UploadFileToStorage(minioClient, context.Background(), username, bucketName, avatarData)
	require.NoError(t, err)
	resData, err := DownloadFileFromStorage(minioClient, context.Background(), username, bucketName)
	require.NoError(t, err)
	require.Equal(t, avatarData, resData)

	err = DeleteFileFromStorage(minioClient, context.Background(), username, bucketName)
	require.NoError(t, err)
}
func GenerateRandomBytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}
