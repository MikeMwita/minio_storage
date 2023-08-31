package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

//func StartServerWithGracefulShutdown(a *fiber.App) {
//	// Create channel for idle connections.
//	idleConnsClosed := make(chan struct{})
//
//	go func() {
//		sigint := make(chan os.Signal, 1)
//		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
//		<-sigint
//
//		// Received an interrupt signal, shutdown.
//		if err := a.Shutdown(); err != nil {
//			// Error from closing listeners, or context timeout:
//			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
//		}
//
//		close(idleConnsClosed)
//	}()
//
//	// Build Fiber connection URL.
//	//fiberConnURL, _ := ConnectionURLBuilder("fiber")
//
//	url := fmt.Sprintf(
//		"%s:%s",
//		"0.0.0.0",
//		"6000",
//	)
//	// Run server.
//	if err := a.Listen(url); err != nil {
//		log.Printf("Oops... Server is not running! Reason: %v", err)
//	}
//
//	<-idleConnsClosed
//}

func StartServer(a *fiber.App) {
	// Build Fiber connection URL.
	//fiberConnURL, _ := ConnectionURLBuilder("fiber")
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
