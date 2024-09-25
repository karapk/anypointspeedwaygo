package main

import (
	"fmt"
	"net/http"
)

type login int
type welcome int

func (l login) ServeHTTP(w http.ResponseWriter, r *http.Request){
	// if r.Method == "GET"{
	// 	fmt.Fprintln(w, "Using GET for login endpoint")
	// }
	if r.Method == "POST"{
		fmt.Fprintln(w, "Using POST for login endpoint")
	}
	switch r.Method{
	case "GET":
		fmt.Fprintln(w, "Using GET for easier login endpoint")
	case "POST":
		fmt.Fprintln(w, "Using POST for easier login endpoint")
	}
	fmt.Fprintln(w, "on log in page")
}

func (wl welcome) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "on Welcome page")
}




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
	//accepts handler accepts patterna dn function signature but not same name
	// http.HandleFunc("/login", mylogin)
	// http.HandleFunc("/welcome", mywelcome)
	// accepts handler interface of type http.Handler
	//accepts pattern and handler
	// http.Handle("/login", http.HandlerFunc(mylogin))
	// http.Handle("/welcome", http.HandlerFunc(mywelcome))
	var i login
	var j welcome
	http.Handle("/login", i)
	http.Handle("/welcome", j)
	fmt.Println("Listening on port 3001")
	// var foo MyServerType
	http.ListenAndServe("localhost:3001", nil)
}
