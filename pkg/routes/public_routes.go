package routes

import (
	"github.com/Filtronic/Minio/app/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PublicRoutes(a *fiber.App, db *gorm.DB) {
	route := a.Group("/api/v1")
	route.Post("/upload", func(c *fiber.Ctx) error {
		return controllers.UploadFile(c, db)
	})
	route.Get("/get-file-metadata", controllers.GetFileMetadata)
	route.Get("/get-file-metadata", controllers.GetBucketMetadata)
	//route.Post("/upload", controllers.UploadFile)
	route.Get("/list-buckets", controllers.ListBuckets)
	route.Delete("/remove-bucket/:bucketName", func(c *fiber.Ctx) error {
		return controllers.RemoveBucket(c, db)
	})
	route.Get("/list-incomplete-uploads/:bucketName", controllers.ListIncompleteUploads)
	route.Post("/set-bucket-tagging/:bucketName", controllers.SetBucketTagging)
	route.Post("/create-bucket/:bucketName", func(c *fiber.Ctx) error {
		return controllers.CreateBucket(c, db)
	})
	route.Get("/get-bucket-tagging/:bucketName", controllers.GetBucketTagging)
	route.Delete("/remove-bucket-tagging/:bucketName", controllers.RemoveBucketTagging)
	route.Post("/set-bucket-replication/:bucketName", controllers.SetBucketReplication)
	route.Post("/enable-versioning/:bucketName", controllers.EnableVersioning)
	route.Get("/get-bucket-replication/:bucketName", controllers.GetBucketReplication)
	route.Post("/remove-bucket-replication/:bucketName", controllers.RemoveBucketReplication)
	route.Get("/get-bucket-replication-metrics/:bucketName", controllers.GetBucketReplicationMetrics)
	route.Get("/get-object/:bucketName/:objectName", controllers.GetObject)
	route.Get("/fget-object/:bucketName/:objectName/:filePath", controllers.FGetObject)
	route.Post("/put-object/:bucketName/:objectName", controllers.PutObject)
	route.Post("/restore-object/:bucketName/:objectName/:versionID", controllers.RestoreObject)
	route.Get("/get-bucket-encryption/:bucketName", controllers.GetBucketEncryption)
	route.Post("/remove-bucket-encryption/:bucketName", controllers.RemoveBucketEncryption)
	route.Post("/set-bucket-policy/:bucketName", controllers.SetBucketPolicy)
	route.Get("/get-bucket-policy/:bucketName", controllers.GetBucketPolicy)

}
