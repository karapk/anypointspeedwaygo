package main

import (
	"fmt"
	"net/http"
)

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

type MyServerType bool

func (m MyServerType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
		<head>	
		</head>
		<body>
			<h1>Hello Worlds</h1>
		</body>
	</html>
	`)
}

func main() {
	var foo MyServerType
	http.ListenAndServe("localhost:3000", foo)
}
