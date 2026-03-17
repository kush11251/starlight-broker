package main

import (
	"github.com/gorilla/mux"
)

func GetEventController(r *mux.Router) {
	r.HandleFunc("/api/events/{id}", getEvent).Methods("GET")
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// logic to get event by id
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}