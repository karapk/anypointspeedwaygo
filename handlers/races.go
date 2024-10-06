package handlers

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
	"sort"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Temperature struct {
    Station     string  `json:"station"`
    Temperature float64 `json:"temperature"`
}

type AverageTemperature struct {
    Station     string  `json:"station"`
    Temperature float64 `json:"temperature"`
}
type Response struct {
    RacerID  string      `json:"racerId"`
    Averages interface{} `json:"averages"`
}

var myDb = make(map[string][]string)

func WelcomeHandler(c echo.Context) error {
    return c.String(http.StatusOK, "Welcome to Anypoint racing! 🚗💨")
}

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

func TemperaturesHandler(c echo.Context) error {
    log.Println("TemperaturesHandler called")
    
    var reader io.ReadCloser
    if c.Request().Header.Get("Content-Encoding") == "gzip" {
        log.Println("Gzip encoding detected")
        gzipReader, err := gzip.NewReader(c.Request().Body)
        if err != nil {
            log.Println("Error creating gzip reader:", err)
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid gzip payload"})
        }
        
        log.Println("Gzip reader successfully created")
        defer func() {
            if err := gzipReader.Close(); err != nil {
                log.Println("Error closing gzip reader:", err)
            } else {
                log.Println("Gzip reader successfully closed")
            }
        }()
        reader = gzipReader
    } else {
        log.Println("No gzip encoding detected, reading raw request body")
        reader = c.Request().Body
    }
    defer func() {
        if err := reader.Close(); err != nil {
            log.Println("Error closing request body reader:", err)
        } else {
            log.Println("Request body reader successfully closed")
        }
    }()

    decoder := json.NewDecoder(reader)
    
    // Log to show start of JSON parsing
    log.Println("Starting to decode JSON payload")

    stationTempSums := make(map[string]float64)
    stationTempCounts := make(map[string]int)

    // Check if the JSON payload starts with an array
    if _, err := decoder.Token(); err != nil {
        log.Println("Error reading JSON start token:", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON payload"})
    }
    log.Println("JSON start token successfully read")

    // Iterate through each JSON object in the array
    for decoder.More() {
        var measurement Temperature
        if err := decoder.Decode(&measurement); err != nil {
            log.Println("Error decoding measurement:", err)
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid measurement"})
        }
        
        log.Printf("Decoded measurement: Station=%s, Temperature=%.2f\n", measurement.Station, measurement.Temperature)
        stationTempSums[measurement.Station] += measurement.Temperature
        stationTempCounts[measurement.Station]++
    }

    // Calculate averages
    log.Println("Calculating average temperatures for each station")
    averages := make([]AverageTemperature, 0, len(stationTempSums))
    for station, sum := range stationTempSums {
        avg := sum / float64(stationTempCounts[station])
        roundedAvg := math.Round(avg*100000) / 100000
        log.Printf("Calculated average for station %s: %.5f\n", station, roundedAvg)
        averages = append(averages, AverageTemperature{
            Station:    station,
            Temperature: roundedAvg,
        })
    }

    // Sort averages alphabetically by station name
    log.Println("Sorting averages by station name")
    sort.Slice(averages, func(i, j int) bool {
        return averages[i].Station < averages[j].Station
    })

    // Prepare and log the response
    response := Response{
        RacerID:  "2532c7d5-511b-466a-a8b7-bb6c797efa36",
        Averages: averages,
    }
    log.Println("Returning response:", response)

    return c.JSON(http.StatusOK, response)
}