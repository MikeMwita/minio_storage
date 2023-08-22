package controllers

import (
	"context"
	"fmt"
	minioUpload "github.com/Filtronic/Minio/platform/minio"
	"github.com/gofiber/fiber/v2"
)

func RemoveBucket(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	err = minioClient.RemoveBucket(ctx, bucketName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("Bucket %s has been successfully removed", bucketName),
	})
}
