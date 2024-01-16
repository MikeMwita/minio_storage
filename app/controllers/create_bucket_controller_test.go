package controllers

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	// Initialize a new SQL mock
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating SQL mock: %v", err)
	}

	// Create a Gorm DB instance using the SQL mock
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDB,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error creating Gorm DB: %v", err)
	}

	return gormDB, mock
}

func TestCreateBucket(t *testing.T) {
	type args struct {
		c  *fiber.Ctx
		db *gorm.DB
	}

	gormDB, _ := setupMockDB(t)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid bucket creation",
			args: args{
				c:  mockFiberCtx("my-bucket"),
				db: gormDB,
			},
			wantErr: false,
		},
		{
			name: "Invalid bucket creation",
			args: args{
				c:  mockFiberCtx("existing-bucket"),
				db: gormDB,
			},
			wantErr: true,
		},
		{
			name: "Empty bucket name",
			args: args{
				c:  mockFiberCtx(""),
				db: gormDB,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateBucket(tt.args.c, tt.args.db)
			assert.Equal(t, (err != nil), tt.wantErr, "CreateBucket() error")
		})
	}
}
