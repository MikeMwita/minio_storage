package controllers

import (
	"context"
	"fmt"
	"github.com/Filtronic/Minio/app/models"
	minioUpload "github.com/Filtronic/Minio/platform/minio"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

func CreateBucket(c *fiber.Ctx, db *gorm.DB) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
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
	bucketTag := "buck01"
	uniqueID := "xyz123"

	// Store metadata in the database
	bucketMetadata := models.BucketMetadata{
		BucketName: bucketName,
		BucketTag:  bucketTag,
		UniqueID:   uniqueID,
	}
	if err := models.CreateBucketMetadata(db, bucketMetadata); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error":       false,
		"bucket_name": bucketName,
		"bucket_tag":  bucketTag,
		"unique_id":   uniqueID,
		"message":     fmt.Sprintf("Bucket %s has been successfully created", bucketName),
	})
}
