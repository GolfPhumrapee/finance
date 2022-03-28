package Minio

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

var (
	minioClient = &minio.Client{}
	ctx         = context.Background()
)

// Client minio client interface
type Client interface {
	UploadFile1(bucketName string, objectName string, filesize int64, filePath io.Reader) error
	UploadFile2(bucketName string, objectName string, filePath string) error
	OsOpenFile(fileHeader *multipart.FileHeader, file multipart.File) (*os.File, error)
}

// Configuration config minio for new connection
type Configuration struct {
	Host            string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

// NewConnection new ftp connection
func NewConnection(config Configuration) (err error) {
	// Initialize minio client object.
	minioClient, err = minio.New(config.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		return err
	}

	return nil
}

type client struct {
	client *minio.Client
}

// GetClient get minio client
func GetClient() Client {
	return &client{
		client: minioClient,
	}
}

const uploadPath = "./upload"

// UploadFile
func (m *client) UploadFile2(bucketName string, objectName string, filePath string) error {
	err := m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		exists, errBucketExists := m.client.BucketExists(ctx, bucketName)
		if errBucketExists != nil {
			logrus.Errorf("[UploadFile] check bucket exists error: %s", err)
			return err
		}

		if !exists {
			logrus.Errorf("[UploadFile] make bucket error: %s", err)
			return err
		}
	}

	_, err = m.client.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		logrus.Errorf("[UploadFile] put object error: %s", err)
		return err
	}

	// object, err := os.Open(objectName)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer object.Close()
	// objectStat, err := object.Stat()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// n, err := m.client.PutObject(ctx, bucketName, filePath, object, objectStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	// log.Println("Uploaded Successfully "+objectName+" of size:: %v", n)
	return nil
}

// UploadFile
func (m *client) UploadFile1(bucketName string, objectName string, filesize int64, filePath io.Reader) error {
	err := m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		exists, errBucketExists := m.client.BucketExists(ctx, bucketName)
		if errBucketExists != nil {
			logrus.Errorf("[UploadFile] check bucket exists error: %s", err)
			return err
		}

		if !exists {
			logrus.Errorf("[UploadFile] make bucket error: %s", err)
			return err
		}
	}

	uploadInfo, err := m.client.PutObject(ctx, bucketName, objectName, filePath, filesize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully uploaded bytes: ", uploadInfo)
	return nil
}

// UploadFile
func (m *client) OsOpenFile(fileHeader *multipart.FileHeader, file multipart.File) (*os.File, error) {

	localFileName := uploadPath + "/" + fileHeader.Filename
	out, err := os.Create(localFileName)
	if err != nil {
		log.Println(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Println(err)
	}
	object, err := os.Open(localFileName)
	return object, err
}
