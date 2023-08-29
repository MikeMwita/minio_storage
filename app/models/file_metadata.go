package models

import (
	"gorm.io/gorm"
	"time"
)

type FileMetadata struct {
	gorm.Model
	FileName   string
	UploaderID int
	UploadTime time.Time
	ETag       string
	VersionID  string
}

func CreateFileMetadata(db *gorm.DB, metadata FileMetadata) error {
	return db.Create(&metadata).Error
}

func GetFileMetadataList(db *gorm.DB) ([]FileMetadata, error) {
	var metadataList []FileMetadata
	result := db.Find(&metadataList)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return metadataList, nil
}
