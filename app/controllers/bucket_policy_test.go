package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestGetBucketPolicy(t *testing.T) {
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
			if err := GetBucketPolicy(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetBucketPolicy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetBucketPolicy(t *testing.T) {
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
				c: mockFiberCtxWithBody("my-bucket", `{"policy":"sample-policy"}`),
			},
			wantErr: false,
		},
		{
			name: "Invalid bucket name",
			args: args{
				c: mockFiberCtxWithBody("non-existent-bucket", `{"policy":"sample-policy"}`),
			},
			wantErr: true,
		},
		{
			name: "Empty bucket name",
			args: args{
				c: mockFiberCtxWithBody("", `{"policy":"sample-policy"}`),
			},
			wantErr: true,
		},
		{
			name: "Invalid request body",
			args: args{
				c: mockFiberCtxWithBody("my-bucket", `invalid-json`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetBucketPolicy(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SetBucketPolicy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func mockFiberCtxWithBody(bucketName, requestBody string) *fiber.Ctx {
	app := fiber.New()

	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.SetMethod("POST")
	ctx.Request.SetRequestURI("/" + bucketName)
	ctx.Request.SetBodyString(requestBody)

	c := app.AcquireCtx(ctx)

	c.Locals("user", context.Background())

	c.Params("bucket_name", bucketName)

	return c
}
