package handlers

import (
    "io"
    "log"
    "net/http"
    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
)

var myDb = make(map[string][]string)

func CreateRaceHandler(c echo.Context) error {
    log.Println("CreateRaceHandler called")

    requestBody := make(map[string]string)
    if err := c.Bind(&requestBody); err != nil {
        log.Println("Error binding request body:", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
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

func CompleteLapHandler(c echo.Context) error {
    log.Println("CompleteLapHandler called")

    // Extract the race ID from the URL parameter
    raceID := c.Param("id")
    log.Println("Received race ID:", raceID)

    // Check if the race exists
    if _, exists := myDb[raceID]; !exists {
        log.Println("Race ID not found:", raceID)
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Race ID not found"})
    }

    // Read the received token from the request body (text/plain)
    receivedTokenBytes, err := io.ReadAll(c.Request().Body)
    if err != nil {
        log.Println("Error reading received token:", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
    }

    receivedToken := string(receivedTokenBytes)
    log.Println("Received token for lap completion:", receivedToken)

    // Retrieve current tokens for the race and append the new token
    tokens := myDb[raceID]
    tokens = append(tokens, receivedToken)
    myDb[raceID] = tokens
    log.Println("Updated tokens for race:", raceID, tokens)

    // Get the previous token
    previousToken := ""
    if len(tokens) > 1 {
        previousToken = tokens[len(tokens)-2]
        log.Println("Previous token found:", previousToken)
    } else {
        log.Println("No valid previous token available")
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "No valid token to return"})
    }

    // Create the response
    response := map[string]string{
        "token":   previousToken,
        "racerId": "2532c7d5-511b-466a-a8b7-bb6c797efa36",
    }
    log.Println("Returning response:", response)

    return c.JSON(http.StatusOK, response)
}

