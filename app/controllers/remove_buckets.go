package controllers

import (
	"context"
	"fmt"
	"github.com/Filtronic/Minio/app/models"
	minioUpload "github.com/Filtronic/Minio/platform/minio"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"os"
)

func RemoveBucket(c *fiber.Ctx, db *gorm.DB) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")

	// Delete bucket from Minio
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	if err := minioClient.RemoveBucket(ctx, bucketName); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	bucketTag := os.Getenv("BUCKET_TAG")
	uniqueID := os.Getenv("UNIQUE_ID")

	// Remove bucket metadata from the database
	if err := models.DeleteBucketMetadata(db, bucketName); err != nil {
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
		"message":     fmt.Sprintf("Bucket %s has been successfully removed", bucketName),
	})
}
