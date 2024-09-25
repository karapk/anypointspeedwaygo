package main

import (
	"net/http"

	"main.go/api"
)

// func main() {
// 	e := echo.New()
// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Welcome to Anypoint racing! 🚗💨")
// 	})

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		e.ServeHTTP(w, r)
// 	})
// 	e.Logger.Print("Listening to port 4000")
// 	e.Logger.Fatal(e.Start(":4000"))
// }

func main() {
	// Register the handler from api package
	http.HandleFunc("/", api.Handler)

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
