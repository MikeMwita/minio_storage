package grpc_handlers

import (
	"context"
	"fmt"
	pb "github.com/Filtronic/Minio/gapi/pb/mutation_gen"
	minioUpload "github.com/Filtronic/Minio/platform/minio"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *grpcServer) RemoveBucket(ctx context.Context, req *pb.RemoveBucketRequest) (*pb.RemoveBucketResponse, error) {
	bucketName := req.GetBucketName()

	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error connecting to Minio: %v", err)
	}
	if err := minioClient.RemoveBucket(ctx, bucketName); err != nil {
		return nil, status.Errorf(codes.Internal, "error removing bucket: %v", err)
	}

	bucketTag := "buck01"
	uniqueID := "xyz123"

	return &pb.RemoveBucketResponse{
		Success:    true,
		BucketName: bucketName,
		BucketTag:  bucketTag,
		UniqueId:   uniqueID,
		Message:    fmt.Sprintf("Bucket %s has been successfully removed", bucketName),
	}, nil
}
