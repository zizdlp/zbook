package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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
	err = CreateBucket(minioClient, context.Background(), "git-files")
	if err != nil {
		return nil, err
	}
	return minioClient, nil
}
func UploadFileToStorage(client *minio.Client, ctx context.Context, objectName string, bucketName string, data []byte) error {
	contentType := http.DetectContentType(data)

	reader := bytes.NewReader(data)

	_, err := client.PutObject(ctx, bucketName, strings.ToLower(objectName), reader, int64(len(data)), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Error().Msgf("Upload object:%s failed: %s", strings.ToLower(objectName), err)
		return err
	}
	return nil
}
func DownloadFileFromStorage(client *minio.Client, ctx context.Context, objectName string, bucketName string) ([]byte, error) {

	// Get object from MinIO
	object, err := client.GetObject(ctx, bucketName, strings.ToLower(objectName), minio.GetObjectOptions{})
	if err != nil {
		log.Error().Msgf("Download object:%s failed: %s", strings.ToLower(objectName), err)
		return nil, err
	}
	defer object.Close()

	// Read object content into a byte slice
	data, err := io.ReadAll(object)
	if err != nil {
		log.Error().Msgf("Read object:%s failed: %s", objectName, err)
		return nil, err
	}
	return data, nil
}
func DeleteFileFromStorage(client *minio.Client, ctx context.Context, objectName string, bucketName string) error {
	//删除一个文件
	err := client.RemoveObject(ctx, bucketName, strings.ToLower(objectName), minio.RemoveObjectOptions{GovernanceBypass: true})
	if err != nil {
		log.Error().Msgf("Delete object: %s failed: %s", strings.ToLower(objectName), err)
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

func ConvertFile2Storage(client *minio.Client, cloneDir string, repoID int64, userID int64, addedFiles []string, modifiedFiles []string, deletedFiles []string) error {
	startTime := time.Now()

	allowedGitFileExtensions := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".svg":  true,
		".gif":  true,
		".webp": true,
	}

	// Helper function to process files (upload or delete)
	processFiles := func(files []string, action func(string) error) error {
		for _, file := range files {
			if err := action(file); err != nil {
				return err
			}
		}
		return nil
	}

	// Upload added and modified files
	uploadFiles := func(file string) error {
		return uploadGitFile(client, context.Background(), cloneDir, file, repoID, userID)
	}

	filteredAddedFiles := util.FilterDiffFilesByExtensions(addedFiles, allowedGitFileExtensions)
	filteredModifiedFiles := util.FilterDiffFilesByExtensions(modifiedFiles, allowedGitFileExtensions)

	if err := processFiles(filteredAddedFiles, uploadFiles); err != nil {
		return fmt.Errorf("failed to upload added files: %w", err)
	}
	if err := processFiles(filteredModifiedFiles, uploadFiles); err != nil {
		return fmt.Errorf("failed to upload modified files: %w", err)
	}

	// Delete files
	deleteFiles := func(file string) error {
		repoIDStr := strconv.FormatInt(repoID, 10)
		name := repoIDStr + "/" + file
		return DeleteFileFromStorage(client, context.Background(), name, "git-files")
	}

	filteredDeletedFiles := util.FilterDiffFilesByExtensions(deletedFiles, allowedGitFileExtensions)

	if err := processFiles(filteredDeletedFiles, deleteFiles); err != nil {
		return fmt.Errorf("failed to delete files: %w", err)
	}

	endTime := time.Now()
	log.Info().Msgf("upload git file to storage: total execution time: %s", endTime.Sub(startTime))

	if _, err := os.Stat(cloneDir); err == nil {
		os.RemoveAll(cloneDir)
	}

	return nil
}
func uploadGitFile(minioClient *minio.Client, ctx context.Context, cloneDir string, filePath string, repoID int64, userID int64) error {

	data, err := os.ReadFile(cloneDir + "/" + filePath)
	if err != nil {
		return err
	}
	ext := strings.ToLower(filePath)
	if strings.HasSuffix(ext, ".png") || strings.HasSuffix(ext, ".jpg") || strings.HasSuffix(ext, ".jpeg") || strings.HasSuffix(ext, ".webp") {
		base64, err := util.ReadImageBytes(cloneDir + "/" + filePath)
		if err != nil {
			return err
		}
		data, err = util.CompressImage(base64)
		if err != nil {
			return err
		}
	}

	userIDStr := strconv.FormatInt(userID, 10)
	repoIDStr := strconv.FormatInt(repoID, 10)
	name := userIDStr + "/" + repoIDStr + "/" + filePath
	err = UploadFileToStorage(minioClient, ctx, name, "git-files", data)
	if err != nil {
		return err
	}

	return nil
}

// DeleteFilesByUserID 删除指定用户 ID 的所有文件
func DeleteAvatarByUsername(client *minio.Client, ctx context.Context, username string) error {
	err := DeleteFileFromStorage(client, ctx, username, "avatar")
	if err != nil {
		log.Error().Msgf("Failed to delete object %s: %s", username, err)
		return err
	}
	log.Info().Msgf("Deleted object %s", username)

	return nil
}

// DeleteFilesByUserID 删除指定用户 ID 的所有文件
func DeleteFilesByUserID(client *minio.Client, ctx context.Context, userID int64, bucketName string) error {
	userIDStr := strconv.FormatInt(userID, 10)

	// 列出用户的所有文件
	objectCh := client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Prefix: userIDStr + "/", Recursive: true})

	// 处理文件删除
	for obj := range objectCh {
		if obj.Err != nil {
			log.Error().Msgf("Failed to list objects: %s", obj.Err)
			return obj.Err
		}

		objectName := obj.Key
		err := DeleteFileFromStorage(client, ctx, objectName, bucketName)
		if err != nil {
			log.Error().Msgf("Failed to delete object %s: %s", objectName, err)
			return err
		}
		log.Info().Msgf("Deleted object %s", objectName)
	}

	return nil
}

// DeleteFilesByUserIDAndRepoID 删除指定 userID 和 repoID 下的所有文件
func DeleteFilesByUserIDAndRepoID(client *minio.Client, ctx context.Context, userID int64, repoID int64, bucketName string) error {
	userIDStr := strconv.FormatInt(userID, 10)
	repoIDStr := strconv.FormatInt(repoID, 10)

	// 构建删除前缀
	prefix := userIDStr + "/" + repoIDStr + "/"

	// 列出指定前缀下的所有文件
	objectCh := client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Prefix: prefix, Recursive: true})

	// 处理文件删除
	for obj := range objectCh {
		if obj.Err != nil {
			log.Error().Msgf("Failed to list objects: %s", obj.Err)
			return obj.Err
		}

		objectName := obj.Key
		err := DeleteFileFromStorage(client, ctx, objectName, bucketName)
		if err != nil {
			log.Error().Msgf("Failed to delete object %s: %s", objectName, err)
			return err
		}
		log.Info().Msgf("Deleted object %s", objectName)
	}

	return nil
}
