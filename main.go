package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"anypointspeedwaygo/handlers"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set the port from the environment or default to 4000
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// Create a new Echo instance
	e := echo.New()

	// Define routes
	e.GET("/", handlers.WelcomeHandler)
	e.POST("/races", handlers.CreateRaceHandler)
	e.POST("/races/:id/laps", handlers.CompleteLapHandler)

	// Start the server
	fmt.Printf("Server running at http://localhost:%s\n", port)
	log.Fatal(e.Start(":" + port))
}
