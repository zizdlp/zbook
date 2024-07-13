package storage

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
	"github.com/zizdlp/zbook/util"
)

func GetMinioClient() (*minio.Client, error) {

	config, err := util.LoadConfig("../")
	if err != nil {
		return nil, err
	}

	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(config.MINIOADDR, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MINIOROOTUSER, config.MINIOROOTPASSWORD, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}
	err = CreateBucket(minioClient, context.Background(), "avatar")
	if err != nil {
		return nil, err
	}
	return minioClient, nil
}
func UploadAvatar(client *minio.Client, ctx context.Context, username string, bucketName string, data []byte) error {
	objectName := username
	contentType := http.DetectContentType(data)

	reader := bytes.NewReader(data)

	_, err := client.PutObject(ctx, bucketName, objectName, reader, int64(len(data)), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Error().Msgf("Upload object failed: %s", err)
		return err
	}
	return nil
}
func DownloadAvatar(client *minio.Client, ctx context.Context, username string, bucketName string) ([]byte, error) {
	objectName := username

	// Get object from MinIO
	object, err := client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Error().Msgf("Download object failed: %s", err)
		return nil, err
	}
	defer object.Close()

	// Read object content into a byte slice
	data, err := io.ReadAll(object)
	if err != nil {
		log.Error().Msgf("Read object failed: %s", err)
		return nil, err
	}
	return data, nil
}
func DeleteAvatar(client *minio.Client, ctx context.Context, username string, bucketName string) error {
	//删除一个文件
	objectName := username
	err := client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{GovernanceBypass: true})
	if err != nil {
		log.Error().Msgf("Delete object failed: %s", err)
		return err
	}
	return nil
}
func CreateBucket(client *minio.Client, ctx context.Context, bucketName string) error {
	err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "cn-south-1", ObjectLocking: false})
	if err != nil {
		exists, errBucketExists := client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Info().Msgf("Bucket %s already exists", bucketName)
			return nil // Bucket already exists, not an error condition
		} else {
			return err // Some other error, return it
		}
	} else {
		log.Info().Msgf("Bucket created successfully: %s", bucketName)
		return nil
	}
}
