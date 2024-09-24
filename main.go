package main

import (
	"fmt"
	"net/http"
)

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

// type MyServerType bool

func myfunc(w http.ResponseWriter, r *http.Request) {
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
	// var foo MyServerType
	http.ListenAndServe("localhost:3001", http.HandlerFunc(myfunc))
}
