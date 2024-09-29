package handler

import (
	"fmt"
	"net/http"

	. "github.com/tbxark/g4vercel"
	"github.com/google/uuid" 
)

var (
	myDb = make(map[string][]string) 
	currentActiveRaceId   string
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

	server.GET("/", func(context *Context) {
		context.String(200, "Welcome to the Race!")
	})

	server.POST("/races", func(context *Context) {
		var requestBody struct {
			Token string `json:"token"`
		}
		
		// Decode the JSON request body
		if err := context.BindJSON(&requestBody); err != nil {
			context.JSON(http.StatusBadRequest, H{"message": "Invalid request"})
			return
		}

		fmt.Println("Received request:", requestBody)
		receivedToken := requestBody.Token
		fmt.Println("Received token:", receivedToken)

		if currentActiveRaceId != "" && myDb[currentActiveRaceId] != nil {
			fmt.Println("Reusing existing race:", currentActiveRaceId)
			context.JSON(http.StatusOK, H{
				"id":      currentActiveRaceId,
				"racerId": "2532c7d5-511b-466a-a8b7-bb6c797efa36",
			})
			return
		}

		raceId := uuid.New().String()
		myDb[raceId] = []string{receivedToken} 
		currentActiveRaceId = raceId         

		toSend := H{
			"id":      raceId,
			"racerId": "2532c7d5-511b-466a-a8b7-bb6c797efa36",
		}

		fmt.Println("Race started:", toSend)
		context.JSON(http.StatusOK, toSend)
	})

	server.Handle(w, r)
}
