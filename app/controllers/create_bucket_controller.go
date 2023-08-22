package controllers

import (
	"context"
	"fmt"
	minioUpload "github.com/Filtronic/Minio/platform/minio"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)

func CreateBucket(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
	// Create minio connection.
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	opts := minio.MakeBucketOptions{
		Region:        "us-east-1",
		ObjectLocking: true,
	}

	err = minioClient.MakeBucket(ctx, bucketName, opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("Bucket %s has been successfully created", bucketName),
	})
}
