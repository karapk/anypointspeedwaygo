package main

import (
	"fmt"
	"net/http"
)

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

// type MyServerType bool

func mylogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
		<head>	
		</head>
		<body>
			<h1>'Log in Anypoint racing--GO! 🚗💨'</h1>
		</body>
	</html>
	`)
}

func mywelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
		<head>	
		</head>
		<body>
			<h1>'Welcome to Anypoint racing--GO! 🚗💨'</h1>
		</body>
	</html>
	`)
}

func main() {
	http.HandleFunc("/login", mylogin)
	http.HandleFunc("/welcome", mywelcome)
	fmt.Println("Listening on port 3001")
	// var foo MyServerType
	http.ListenAndServe("localhost:3001", nil)
}
