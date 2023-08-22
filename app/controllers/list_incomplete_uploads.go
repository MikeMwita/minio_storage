package controllers

import (
	"context"
	minioUpload "github.com/Filtronic/Minio/platform/minio"
	"github.com/gofiber/fiber/v2"
)

type MultipartObjectResponse struct {
	Key      string `json:"key"`
	UploadID string `json:"upload_id"`
	Size     int64  `json:"size"`
}

func ListIncompleteUploads(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	prefix := c.Query("prefix")
	recursive := c.Query("recursive") == "true"

	var multipartObjectList []MultipartObjectResponse

	multipartObjectCh := minioClient.ListIncompleteUploads(ctx, bucketName, prefix, recursive)
	for multiPartObject := range multipartObjectCh {
		if multiPartObject.Err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   multiPartObject.Err.Error(),
			})
		}
		multipartObjectList = append(multipartObjectList, MultipartObjectResponse{
			Key:      multiPartObject.Key,
			UploadID: multiPartObject.UploadID,
			Size:     multiPartObject.Size,
		})
	}
	return c.JSON(fiber.Map{
		"error":             false,
		"multipart_objects": multipartObjectList,
	})
}
