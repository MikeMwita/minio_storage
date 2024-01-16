package controllers

import "testing"

func TestGetBucketTagging(t *testing.T) {
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetBucketTagging(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SetBucketTagging() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
