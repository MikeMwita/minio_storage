package models

import "time"

type BucketResponse struct {
	Name         string    `json:"name"`
	CreationDate time.Time `json:"creation_date"`
}
