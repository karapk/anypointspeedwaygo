package handler

import (
	"net/http"
	"encoding/json"
	."github.com/tbxark/g4vercel"
	"github.com/google/uuid"
)


var myDb = struct {
	races map[string][]string
}{
	races: make(map[string][]string),
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

\
func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Anypoint Racing! 🏎️"))
		return
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}


func handlePost(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/races":
		startRace(w, r)
	case "/races/":
		
		id := r.URL.Path[len("/races/"):]
		completeLap(w, r, id)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}


func startRace(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Token string `json:"token"`
	}

	
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	raceId := uuid.New().String()
	myDb.races[raceId] = []string{body.Token} // Initialize with the starting token

	toSend := map[string]interface{}{
		"id":      raceId,
		"racerId": "2532c7d5-511b-466a-a8b7-bb6c797efa36",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(toSend)
}


func completeLap(w http.ResponseWriter, r *http.Request, raceId string) {
	var receivedToken string

	
	if err := json.NewDecoder(r.Body).Decode(&receivedToken); err != nil {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}

	tokens, ok := myDb.races[raceId]
	if !ok {
		http.Error(w, "Race ID not found", http.StatusNotFound)
		return
	}

	tokens = append(tokens, receivedToken) 
	myDb.races[raceId] = tokens

	if len(tokens) < 2 {
		http.Error(w, "No valid token to return", http.StatusBadRequest)
		return
	}

	previousToken := tokens[len(tokens)-2]
	toSend := map[string]interface{}{
		"token":   previousToken,
		"racerId": "2532c7d5-511b-466a-a8b7-bb6c797efa36",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(toSend)
}
