package models

import (
	"gorm.io/gorm"
	"time"
)

type BucketResponse struct {
	Name         string    `json:"name"`
	CreationDate time.Time `json:"creation_date"`
}

type BucketMetadata struct {
	gorm.Model
	BucketName string
	BucketTag  string
	UniqueID   string
}

func CreateBucketMetadata(db *gorm.DB, metadata BucketMetadata) error {
	return db.Create(&metadata).Error
}

func GetBucketMetadataList(db *gorm.DB) ([]BucketMetadata, error) {
	var metadataList []BucketMetadata
	result := db.Find(&metadataList)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return metadataList, nil
}

func DeleteBucketMetadata(db *gorm.DB, bucketName string) error {
	return db.Where("bucket_name = ?", bucketName).Delete(&BucketMetadata{}).Error
}
