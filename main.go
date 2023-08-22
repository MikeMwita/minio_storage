package main

import (
	"github.com/Filtronic/Minio/pkg/configs"
	"github.com/Filtronic/Minio/pkg/routes"
	"github.com/Filtronic/Minio/pkg/utils"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	config := configs.FiberConfig()
	// Define a new Fiber app with config.
	app := fiber.New(config)
	routes.PublicRoutes(app)
	utils.StartServer(app)

	//err := app.Listen(":6000")
	//if err != nil {
	//	log.Fatalf("Failed to start server: %v", err)
}
