package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/events", getEvents).Methods("GET")
	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]string{"event1", "event2"})
}