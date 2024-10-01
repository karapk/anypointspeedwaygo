package main

import (
	"fmt"
	"log"
	"os"

	// "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"anypointspeedwaygo/handlers"
)

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	e := echo.New()

	// Define routes
	e.GET("/", handlers.WelcomeHandler)
	e.POST("/races", handlers.CreateRaceHandler)
	e.POST("/races/:id/laps", handlers.CompleteLapHandler)

	// Start the server
	fmt.Printf("Server running at http://localhost:%s\n", port)
	log.Fatal(e.Start(":" + port))
}
