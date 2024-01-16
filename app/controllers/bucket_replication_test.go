package controllers

import (
	"github.com/gofiber/fiber/v2"
	"testing"
)

func TestGetBucketReplication(t *testing.T) {
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
			if err := GetBucketReplication(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetBucketReplication() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetBucketReplicationMetrics(t *testing.T) {
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
			if err := GetBucketReplicationMetrics(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetBucketReplicationMetrics() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRemoveBucketReplication(t *testing.T) {
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
			if err := RemoveBucketReplication(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("RemoveBucketReplication() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetBucketReplication(t *testing.T) {
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
			if err := SetBucketReplication(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SetBucketReplication() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
