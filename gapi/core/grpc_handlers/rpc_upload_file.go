package grpc_handlers

import (
	"bytes"
	"context"
	"github.com/Filtronic/Minio/app/models"
	pb "github.com/Filtronic/Minio/gapi/pb/mutation_gen"
	minioUpload "github.com/Filtronic/Minio/platform/minio"
	"github.com/minio/minio-go/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"time"
)

func (server *grpcServer) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	bucketName := "filtronic"
	fileContent := req.GetFileContent()
	//objectName := fileContent.GetFileName()
	objectName := "testfile.txt"
	contentType := fileContent.GetContentType()
	fileSize := int64(len(fileContent.GetContent()))

	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error connecting to Minio: %v", err)
	}
	fileBuffer := ioutil.NopCloser(bytes.NewReader(fileContent.GetContent()))

	info, err := minioClient.PutObject(ctx, bucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error uploading file: %v", err)
	}

	// Store metadata in the database
	fileMetadata := models.FileMetadata{
		FileName:   objectName,
		UploaderID: 123,
		UploadTime: time.Now(),
		ETag:       info.ETag,
		VersionID:  info.VersionID,
	}
	if err := server.db.Create(&fileMetadata).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "error storing file metadata: %v", err)
	}

	return &pb.UploadFileResponse{
		Success: true,
		Message: "File uploaded successfully.",
		//Info: &pb.UploadFileInfo{
		//	Etag:      info.ETag,
		//	VersionId: info.VersionID,
		//},
	}, nil
}
