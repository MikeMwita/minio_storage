package controllers

import "testing"

func TestListBuckets(t *testing.T) {
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
			if err := ListBuckets(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("ListBuckets() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
