package controllers

import (
	"github.com/gofiber/fiber/v2"
	"testing"
)

func TestGetBucketTagging(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid bucket name",
			args: args{
				c: mockFiberCtx("my-bucket"),
			},
			wantErr: false,
		},
		{
			name: "Invalid bucket name",
			args: args{
				c: mockFiberCtx("non-existent-bucket"),
			},
			wantErr: true,
		},
		{
			name: "Empty bucket name",
			args: args{
				c: mockFiberCtx(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetBucketTagging(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetBucketTagging() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRemoveBucketTagging(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid bucket name",
			args: args{
				c: mockFiberCtx("my-bucket"),
			},
			wantErr: false,
		},
		{
			name: "Invalid bucket name",
			args: args{
				c: mockFiberCtx("non-existent-bucket"),
			},
			wantErr: true,
		},
		{
			name: "Empty bucket name",
			args: args{
				c: mockFiberCtx(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveBucketTagging(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("RemoveBucketTagging() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetBucketTagging(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid bucket name",
			args: args{
				c: mockFiberCtx("my-bucket"),
			},
			wantErr: false,
		},
		{
			name: "Invalid bucket name",
			args: args{
				c: mockFiberCtx("non-existent-bucket"),
			},
			wantErr: true,
		},
		{
			name: "Empty bucket name",
			args: args{
				c: mockFiberCtx(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetBucketTagging(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SetBucketTagging() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
