package grpc_handlers

import (
	"context"
	"fmt"
	pb "github.com/Filtronic/Minio/gapi/pb/mutation_gen"
	minioUpload "github.com/Filtronic/Minio/platform/minio"
	"github.com/minio/minio-go/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *grpcServer) CreateBucket(ctx context.Context, req *pb.CreateBucketRequest) (*pb.CreateBucketResponse, error) {
	bucketName := req.GetBucketName()
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error connecting to Minio: %v", err)
	}
	opts := minio.MakeBucketOptions{
		Region:        "us-east-1",
		ObjectLocking: true,
	}

	err = minioClient.MakeBucket(ctx, bucketName, opts)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating bucket: %v", err)
	}

	bucketTag := "buck01"
	uniqueID := "xyz123"

	return &pb.CreateBucketResponse{
		Success:    true,
		BucketName: bucketName,
		BucketTag:  bucketTag,
		UniqueId:   uniqueID,
		Message:    fmt.Sprintf("Bucket %s has been successfully created", bucketName),
	}, nil
}
