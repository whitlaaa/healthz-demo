package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var started = time.Now()

func main() {
	log.Println("Sample healthz app started...")

	r := mux.NewRouter()
	r.HandleFunc("/healthz", healthHandler)
	r.HandleFunc("/readyz", readyHandler)
	r.HandleFunc("/unready", unreadyHandler)
	r.HandleFunc("/demo", demoHandler)

	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking liveness...")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(buildResponse("Healthy!"))
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking readiness...")
	var response Response
	elapsed := time.Now().Sub(started)
	if elapsed.Seconds() > 20 {
		w.WriteHeader(http.StatusOK)
		response = buildResponse("Ready!")
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		response = buildResponse("Not ready... :(")
	}
	json.NewEncoder(w).Encode(response)
}

func unreadyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Resetting readiness...")
	started = time.Now()
	w.WriteHeader(http.StatusOK)
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Demo endpoint...")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(buildResponse("The endpoint works!"))
}

func buildResponse(message string) Response {
	return Response{time.Now(), message}
}

type Response struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}
