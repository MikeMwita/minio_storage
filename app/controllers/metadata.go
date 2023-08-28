package controllers

import (
	"github.com/Filtronic/Minio/app/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

//	func GetFileMetadata(c *fiber.Ctx) error {
//		db := c.Locals("db").(*gorm.DB)
//		// Fetch metadata from the database
//		metadataList, err := models.GetFileMetadataList(db)
//		if err != nil {
//			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//				"error": true,
//				"msg":   err.Error(),
//			})
//		}
//		return c.JSON(fiber.Map{
//			"error":    false,
//			"metadata": metadataList,
//		})
//	}

func GetFileMetadata(c *fiber.Ctx) error {
	// Get the 'db' instance from the fiber context
	db := c.Locals("db").(*gorm.DB)
	if db == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Database connection is not found",
		})
	}
	// Fetch metadata from the database
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
