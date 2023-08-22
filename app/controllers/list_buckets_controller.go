package controllers

import (
	"context"
	minioUpload "github.com/Filtronic/Minio/platform/minio"
	"github.com/gofiber/fiber/v2"
	"time"
)

type BucketResponse struct {
	Name         string    `json:"name"`
	CreationDate time.Time `json:"creation_date"`
}

func ListBuckets(c *fiber.Ctx) error {
	ctx := context.Background()
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	buckets, err := minioClient.ListBuckets(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	var bucketList []BucketResponse
	for _, bucket := range buckets {
		bucketList = append(bucketList, BucketResponse{
			Name:         bucket.Name,
			CreationDate: bucket.CreationDate,
		})
	}
	return c.JSON(fiber.Map{
		"error":   false,
		"buckets": bucketList,
	})
}
