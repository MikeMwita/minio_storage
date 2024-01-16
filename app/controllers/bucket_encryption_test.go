package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestGetBucketEncryption(t *testing.T) {
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
			if err := GetBucketEncryption(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetBucketEncryption() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRemoveBucketEncryption(t *testing.T) {
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
			if err := RemoveBucketEncryption(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("RemoveBucketEncryption() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// mockFiberCtx creates a mock fiber context with the given bucket name as a parameter
func mockFiberCtx(bucketName string) *fiber.Ctx {
	app := fiber.New()

	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.SetMethod("GET")
	ctx.Request.SetRequestURI("/" + bucketName)

	c := app.AcquireCtx(ctx)

	c.Locals("user", context.Background())

	c.Params("mike", bucketName)

	return c
}
