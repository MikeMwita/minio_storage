package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func StartServer(a *fiber.App) {
	// Build Fiber connection URL.
	url := fmt.Sprintf(
		"%s:%s",
		"0.0.0.0",
		"6000",
	)
	// Run server.
	if err := a.Listen(url); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
