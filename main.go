package main

import (
	"fmt"
	"log"
	"os"

	"anypointspeedwaygo/handlers"
	"github.com/labstack/echo/v4"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	e := echo.New()

	// Define routes
	e.GET("/", handlers.WelcomeHandler)
	e.POST("/races", handlers.CreateRaceHandler)
	e.POST("/races/:id/laps", handlers.CompleteLapHandler)
	e.POST("/temperatures", handlers.TemperaturesHandler)

	fmt.Printf("Server running at http://localhost:%s\n", port)
	log.Fatal(e.Start(":" + port))
}
