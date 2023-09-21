package controllers

import (
	"context"
	"fmt"
	minioUpload "github.com/Filtronic/Minio/platform/minio"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7/pkg/tags"
)

// set bucket tagging

func SetBucketTagging(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	tagMap := map[string]string{
		"Tag1": "Value1",
		"Tag2": "Value2",
	}
	bucketTags, err := tags.NewTags(tagMap, false)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err = minioClient.SetBucketTagging(ctx, bucketName, bucketTags)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("Tags set successfully for bucket %s", bucketName),
	})
}

//get bucket tagging

func GetBucketTagging(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	tags, err := minioClient.GetBucketTagging(ctx, bucketName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"tags":  tags,
	})
}

// Remove bucket tagging

func RemoveBucketTagging(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	err = minioClient.RemoveBucketTagging(ctx, bucketName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("Tags removed successfully from bucket %s", bucketName),
	})
}
