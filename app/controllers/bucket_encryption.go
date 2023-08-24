package controllers

import (
	"context"
	"fmt"
	minioUpload "github.com/Filtronic/Minio/platform/minio"
	"github.com/gofiber/fiber/v2"
)

// get bucket encryption

func GetBucketEncryption(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Get default encryption configuration
	encryptionConfig, err := minioClient.GetBucketEncryption(ctx, bucketName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error":  false,
		"config": encryptionConfig,
		"msg":    "Default encryption configuration retrieved successfully.",
	})
}

// remove bucket encryption

func RemoveBucketEncryption(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	err = minioClient.RemoveBucketEncryption(ctx, bucketName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   fmt.Sprintf("Default encryption configuration removed from bucket '%s'.", bucketName),
	})
}
