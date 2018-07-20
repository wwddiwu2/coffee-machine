package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mkqavi/coffee-machine/v0/pkg/coffeemachine"
)

var m coffeemachine.Machine

func main() {
	m = coffeemachine.New()

	r := mux.NewRouter()
	r.HandleFunc("/", machineStatus).Methods("GET")
	r.HandleFunc("/clean", clean).Methods("POST")
	r.HandleFunc("/brew", brew).Methods("POST")

	log.Fatal(http.ListenAndServe(":31565", r))
}

func machineStatus(w http.ResponseWriter, r *http.Request) {
	contentJSON(w)

	err := json.NewEncoder(w).Encode(status{
		Cleanliness: m.Cleanliness(),
		Status:      uint8(m.Status()),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Fatal(err)
	}
}

func clean(w http.ResponseWriter, r *http.Request) {
	contentJSON(w)

	enc := json.NewEncoder(w)
	err := m.Clean()
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		sendError(enc, err)
	}
	w.Write([]byte("{}"))
}

func brew(w http.ResponseWriter, r *http.Request) {
	contentJSON(w)

	var c coffee
	enc := json.NewEncoder(w)

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		sendError(enc, err)
	}

	err = m.Pour(coffeemachine.Drink(c.Coffee))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		sendError(enc, err)
	}

	w.Write([]byte("{}"))
}

func sendError(e *json.Encoder, err error) {
	e.Encode(apiError{err.Error()})
}

func contentJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
