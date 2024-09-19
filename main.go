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
	fmt.Fprintln(w, "Hello World")
	fmt.Fprintf(w, "Request is : %+v", r)
}

func main() {
	var foo MyServerType
	http.ListenAndServe("localhost:3000", foo)
}
