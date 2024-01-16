package controllers

import (
	"context"
	"github.com/Filtronic/Minio/app/models"
	minioUpload "github.com/Filtronic/Minio/platform/minio"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"log"
	"time"
)

func UploadFile(c *fiber.Ctx, db *gorm.DB) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
	file, err := c.FormFile("fileUpload")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	buffer, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	defer buffer.Close()
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	objectName := file.Filename
	fileBuffer := buffer
	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size
	info, err := minioClient.PutObject(ctx, bucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	// Store metadata in the database
	fileMetadata := models.FileMetadata{
		FileName:   objectName,
		UploaderID: 123,
		UploadTime: time.Now(),
		ETag:       info.ETag,
		VersionID:  info.VersionID,
	}
	if err := models.CreateFileMetadata(db, fileMetadata); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "File uploaded successfully.",
		"info":  info,
	})
}
