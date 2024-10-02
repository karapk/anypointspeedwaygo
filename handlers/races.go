package handlers

import (
    "io"
    "net/http"
    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
)

var myDb = make(map[string][]string)

func CreateRaceHandler(c echo.Context) error {

    requestBody := make(map[string]string)
    if err := c.Bind(&requestBody); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
    }

    token := requestBody["token"]

    raceID := uuid.New().String()

    myDb[raceID] = []string{token}

    response := map[string]string{
        "id":      raceID,
        "racerId": "2532c7d5-511b-466a-a8b7-bb6c797efa36",
    }
    return c.JSON(http.StatusOK, response)
}

func CompleteLapHandler(c echo.Context) error {

    raceID := c.Param("id")

    if _, exists := myDb[raceID]; !exists {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Race ID not found"})
    }
    receivedTokenBytes, err := io.ReadAll(c.Request().Body)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
    }

    receivedToken := string(receivedTokenBytes)

    tokens := myDb[raceID]
    tokens = append(tokens, receivedToken)
    myDb[raceID] = tokens

    previousToken := ""
    if len(tokens) > 1 {
        previousToken = tokens[len(tokens)-2]
    } else {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "No valid token to return"})
    }

    response := map[string]string{
        "token":   previousToken,
        "racerId": "2532c7d5-511b-466a-a8b7-bb6c797efa36",
    }

    return c.JSON(http.StatusOK, response)
}

