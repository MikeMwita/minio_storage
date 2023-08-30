package minioUpload

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func MinioConnection() (*minio.Client, error) {
	//ctx := context.Background()
	s := "s3-api.filtronic.co.ke"
	ak := "Vd2DlMhoe1W4QKbSQDLw"
	ask := "FTaNk44w4GQIwo2jdv6zyGUL5KiDUHvj9LVyroE3"
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
