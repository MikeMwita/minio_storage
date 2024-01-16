package controllers

import "testing"

func TestGetBucketReplication(t *testing.T) {
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetBucketReplication(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SetBucketReplication() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
