package controllers

import (
	"github.com/Filtronic/Minio/app/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetFileMetadata(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	if db == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Database connection is not found",
		})
	}
	metadataList, err := models.GetFileMetadataList(db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error":    false,
		"metadata": metadataList,
	})
}

func GetBucketMetadata(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	if db == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Database connection is not found",
		})
	}
	metadataList, err := models.GetBucketMetadataList(db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error":    false,
		"metadata": metadataList,
	})
}
