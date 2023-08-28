//package main
//
//import (
//	"github.com/Filtronic/Minio/pkg/configs"
//	"github.com/Filtronic/Minio/pkg/routes"
//	"github.com/Filtronic/Minio/pkg/utils"
//	"github.com/gofiber/fiber/v2"
//	_ "github.com/joho/godotenv/autoload"
//)
//
//func main() {
//
//	config := configs.FiberConfig()
//	// Define a new Fiber app with config.
//	app := fiber.New(config)
//	routes.PublicRoutes(app)
//	utils.StartServer(app)
//
//}

package main

import (
	"github.com/Filtronic/Minio/app/models"
	"github.com/Filtronic/Minio/pkg/configs"
	"github.com/Filtronic/Minio/pkg/routes"
	"github.com/Filtronic/Minio/pkg/utils"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)
	db = initDB()
	routes.PublicRoutes(app, db)
	utils.StartServer(app)
}

func initDB() *gorm.DB {
	dsn := "host=localhost user=filtronic dbname=edms sslmode=disable password=secret"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	db.AutoMigrate(&models.FileMetadata{})

	return db
}
