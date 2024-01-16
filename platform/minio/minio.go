package minioUpload

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
)

func MinioConnection() (*minio.Client, error) {
	s := os.Getenv("MINIO_HOST") + ":" + os.Getenv("MINIO_PORT")
	ak := os.Getenv("MINIO_ACCESSKEY")
	ask := "MINIO_ACCESSKEY_SECRET"
	useSSL := true
	minioClient, errInit := minio.New(s, &minio.Options{
		Creds:  credentials.NewStaticV4(ak, ask, ""),
		Secure: useSSL,
	})
	if errInit != nil {
		log.Fatalln(errInit)
	}

	return minioClient, errInit
}
