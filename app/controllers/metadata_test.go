package controllers

import (
	"github.com/gofiber/fiber/v2"
	"testing"
)

func TestGetBucketMetadata(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetBucketMetadata(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetBucketMetadata() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFileMetadata(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetFileMetadata(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetFileMetadata() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
