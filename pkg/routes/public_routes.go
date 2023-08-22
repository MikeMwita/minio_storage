package routes

import (
	"github.com/Filtronic/Minio/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/upload", controllers.UploadFile)
	route.Get("/list-buckets", controllers.ListBuckets)
	route.Delete("/remove-bucket/:bucketName", controllers.RemoveBucket)
	route.Get("/list-incomplete-uploads/:bucketName", controllers.ListIncompleteUploads)
	route.Post("/set-bucket-tagging/:bucketName", controllers.SetBucketTagging)
	route.Post("/create-bucket/:bucketName", controllers.CreateBucket)
	route.Get("/get-bucket-tagging/:bucketName", controllers.GetBucketTagging)
	route.Delete("/remove-bucket-tagging/:bucketName", controllers.RemoveBucketTagging)

}
