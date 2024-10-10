package handlers

import (
    "compress/gzip"
    "encoding/json"
    "io"
    "log"
    "math"
    "net/http"
    "sort"

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

func TemperaturesHandler(c echo.Context) error {
    log.Println("TemperaturesHandler called")

    log.Println("Content-Length:", c.Request().Header.Get("Content-Length"))

    var reader io.ReadCloser
    log.Println("Checking for gzip encoding")
    
    if c.Request().Header.Get("Content-Encoding") == "gzip" {
        log.Println("Gzip encoding detected")
        gzipReader, err := gzip.NewReader(c.Request().Body)
        if err != nil {
            log.Println("Error creating gzip reader:", err)
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid gzip payload"})
        }

        log.Println("Gzip reader successfully created")
        defer func() {
            log.Println("Closing gzip reader")
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
        log.Println("Closing request body reader")
        if err := reader.Close(); err != nil {
            log.Println("Error closing request body reader:", err)
        } else {
            log.Println("Request body reader successfully closed")
        }
    }()

    decoder := json.NewDecoder(reader)

    // Log to show start of JSON parsing
    log.Println("Starting to decode JSON payload")

    stationTempSums := make(map[string]*float64)
    stationTempCounts := make(map[string]*int)

    // Check if the JSON payload starts with an array
    if _, err := decoder.Token(); err != nil {
        log.Println("Error reading JSON start token:", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON payload"})
    }
    log.Println("JSON start token successfully read")

    // Initialize a counter for the number of measurements processed
    var measurementCount int

    // Iterate through each JSON object in the array
    for decoder.More() {
        var measurement Temperature
        if err := decoder.Decode(&measurement); err != nil {
            log.Println("Error decoding measurement:", err)
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid measurement"})
        }

        // If the station hasn't been seen yet, initialize it
        if _, exists := stationTempSums[measurement.Station]; !exists {
            tempSum := 0.0
            count := 0
            stationTempSums[measurement.Station] = &tempSum
            stationTempCounts[measurement.Station] = &count
        }

        
        *stationTempSums[measurement.Station] += measurement.Temperature
        *stationTempCounts[measurement.Station]++

        
        measurementCount++

        if measurementCount%10000000 == 0 {
            log.Printf("Processed %d measurements so far", measurementCount)
        }
    }

    
    log.Println("Calculating average temperatures for each station")
    averages := make([]AverageTemperature, 0, len(stationTempSums))
    for station, sumPtr := range stationTempSums {
        avg := *sumPtr / float64(*stationTempCounts[station])
        roundedAvg := math.Round(avg*100000) / 100000
        averages = append(averages, AverageTemperature{
            Station:     station,
            Temperature: roundedAvg,
        })
    }


    log.Println("Sorting averages by station name")
    sort.Slice(averages, func(i, j int) bool {
        return averages[i].Station < averages[j].Station
    })

    
    response := Response{
        RacerID:  "2532c7d5-511b-466a-a8b7-bb6c797efa36",
        Averages: averages,
    }
    log.Println("Returning response:", response)

    return c.JSON(http.StatusOK, response)
}
