package controllers

import "testing"

func TestGetBucketPolicy(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetBucketPolicy(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SetBucketPolicy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
