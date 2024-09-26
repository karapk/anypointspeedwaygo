package handler

import (
	"fmt"
	"net/http"

	. "github.com/tbxark/g4vercel"
	"github.com/google/uuid" 
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()
	server.Use(Recovery(func(err interface{}, c *Context) {
		if httpError, ok := err.(HttpError); ok {
			c.JSON(httpError.Status, H{
				"message": httpError.Error(),
			})
		} else {
			message := fmt.Sprintf("%s", err)
			c.JSON(500, H{
				"message": message,
			})
		}
	}))

	// Welcome endpoint
	server.GET("/", func(context *Context) {
		context.String(200, "Welcome to Anypoint Racing! 🚗💨")
	})

	// Start a new race
	server.POST("/races", func(context *Context) {
		var body struct {
			Token string `json:"token"`
		}
		if err := context.Bind(&body); err != nil {
			context.JSON(400, H{"message": "Invalid request"})
			return
		}

		raceID := uuid.New().String() // Generate a new unique race ID
		// Here you would typically store the raceID and token in a database or in-memory store

		context.JSON(200, H{
			"id":       raceID,
			"racerId":  "2532c7d5-511b-466a-a8b7-bb6c797efa36",
		})
	})

	// Handle lap completion
	server.POST("/races/:id/laps", func(context *Context) {
		raceID := context.Param("id")
		var body struct {
			Token string `json:"token"`
		}
		if err := context.Bind(&body); err != nil {
			context.JSON(400, H{"message": "Invalid request"})
			return
		}

		// Logic to handle lap completion would go here
		// For demonstration, we'll just respond with the received token

		context.JSON(200, H{
			"token":   body.Token,
			"racerId": "2532c7d5-511b-466a-a8b7-bb6c797efa36",
		})
	})

	server.Handle(w, r)
}
