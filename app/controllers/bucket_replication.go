package controllers

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/Filtronic/Minio/platform/minio"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7/pkg/replication"
)

func SetBucketReplication(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Replication configuration string
	replicationStr := `
	<ReplicationConfiguration>
		<Role>arn:minio:s3::598361bf-3cec-49a7-b529-ce870a34d759:*</Role>
		<Rule>
			<DeleteMarkerReplication>
				<Status>Disabled</Status>
			</DeleteMarkerReplication>
			<Destination>
				<Bucket>destination-bucket-name</Bucket>
				<StorageClass>STANDARD</StorageClass>
			</Destination>
			<Filter>
				<And>
					<Prefix>string</Prefix>
					<Tag>
						<Key>string</Key>
						<Value>string</Value>
					</Tag>
				</And>
				<Prefix>string</Prefix>
				<Tag>
					<Key>string</Key>
					<Value>string</Value>
				</Tag>
			</Filter>
			<ID>string</ID>
			<Prefix>string</Prefix>
			<Priority>1</Priority>
			<Status>Enabled</Status>
		</Rule>
	</ReplicationConfiguration>
	`
	replicationConfig := replication.Config{}
	if err := xml.Unmarshal([]byte(replicationStr), &replicationConfig); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err = minioClient.SetBucketReplication(ctx, bucketName, replicationConfig)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("Replication configuration set successfully for bucket %s", bucketName),
	})
}

//get bucket replication

func GetBucketReplication(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	replicationConfig, err := minioClient.GetBucketReplication(ctx, bucketName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":             false,
		"replicationConfig": replicationConfig,
	})
}

//Remove bucket Replication

func RemoveBucketReplication(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName") // Assuming the bucket name is passed as a parameter
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err = minioClient.RemoveBucketReplication(ctx, bucketName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("Replication configuration removed for bucket %s", bucketName),
	})
}

// get bucket replication metrics
func GetBucketReplicationMetrics(c *fiber.Ctx) error {
	ctx := context.Background()
	bucketName := c.Params("bucketName")
	minioClient, err := minioUpload.MinioConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	replicationMetrics, err := minioClient.GetBucketReplicationMetrics(ctx, bucketName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error":              false,
		"replicationMetrics": replicationMetrics,
	})
}
