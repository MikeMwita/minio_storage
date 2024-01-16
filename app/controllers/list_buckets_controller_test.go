package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type MockMinioClient struct{}

func (m *MockMinioClient) ListBuckets(ctx context.Context) ([]minio.BucketInfo, error) {
	// Mock data for testing
	bucketList := []minio.BucketInfo{
		{Name: "bucket1", CreationDate: time.Now()},
		{Name: "bucket2", CreationDate: time.Now().Add(-24 * time.Hour)},
	}

	return bucketList, nil
}

func TestListBuckets(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "List buckets successfully",
			args: args{
				c: &fiber.Ctx{},
			},
			wantErr: false,
		},
	}

	//// Override the MinioConnection function with the mock Minio client
	//minioUpload.MinioConnection = func() (minioUpload.MinioClient, error) {
	//	return &MockMinioClient{}, nil
	//}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ListBuckets(tt.args.c)
			assert.Equal(t, (err != nil), tt.wantErr, "ListBuckets() error")
		})
	}
}
