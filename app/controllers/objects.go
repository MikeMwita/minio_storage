package controllers

import (
	"context"
	"fmt"
	"github.com/Filtronic/Minio/platform/minio"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"io"
	"os"
)

func GetObject(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
	objectName := c.Params("objectName")
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	object, err := minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	defer object.Close()
	localFilePath := fmt.Sprintf("/tmp/%s", objectName)
	localFile, err := os.Create(localFilePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	defer localFile.Close()

	if _, err := io.Copy(localFile, object); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("Object '%s' downloaded to '%s'", objectName, localFilePath),
	})
}

//Download and save the object as a file in the local filesystem

func FGetObject(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
	objectName := c.Params("objectName")
	filePath := c.Params("filePath")

	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err = minioClient.FGetObject(ctx, bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("Object '%s' downloaded and saved to '%s'", objectName, filePath),
	})
}

// upload objects < 128MiB in a single PUT operation

func PutObject(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
	objectName := c.Params("objectName")
	// Read the file to upload
	file, err := c.FormFile("fileUpload")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	fileReader, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	defer fileReader.Close()

	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	uploadInfo, err := minioClient.PutObject(
		ctx,
		bucketName,
		objectName,
		fileReader,
		file.Size,
		minio.PutObjectOptions{
			ContentType: file.Header["Content-Type"][0],
		},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"info":  uploadInfo,
	})
}

//restore object

func RestoreObject(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
	objectName := c.Params("objectName")
	versionID := c.Params("versionID")

	// Create a RestoreRequest with desired options
	opts := minio.RestoreRequest{}
	opts.SetDays(1)
	opts.SetGlacierJobParameters(minio.GlacierJobParameters{Tier: minio.TierStandard})
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Restore the object
	err = minioClient.RestoreObject(
		ctx,
		bucketName,
		objectName,
		versionID,
		opts,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "Object restoration initiated successfully.",
	})
}
