package handlers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)


func CreateRaceHandler(c echo.Context) error {
    log.Println("CreateRaceHandler called")

    requestBody := make(map[string]string)
    if err := c.Bind(&requestBody); err != nil {
        log.Println("Error binding request body:", err)
        return c.JSON(
        	http.StatusBadRequest,
        	map[string]string{"error": "Invalid request payload"},
        )
    }

    // Extract the token from the request body
    token := requestBody["token"]
    log.Println("Received token:", token)

    // Generate a new race ID
    raceID := uuid.New().String()
    log.Println("Generated race ID:", raceID)

    // Store the token in the database
    myDb[raceID] = []string{token}
    log.Println("Race ID and token saved to myDb:", raceID, myDb[raceID])

    // Create the response
    response := map[string]string{
        "id":      raceID,
        "racerId": "2532c7d5-511b-466a-a8b7-bb6c797efa36",
    }
    log.Println("Returning response:", response)

    return c.JSON(http.StatusOK, response)
}