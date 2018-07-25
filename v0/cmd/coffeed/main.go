package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/mkqavi/coffee-machine/v0/pkg/coffeemachine"
)

var (
	m      coffeemachine.Machine
	logger *log.Logger
)

func main() {
	stop := make(chan os.Signal, 1)
	<-serveAPI(stop)
}

func serveAPI(stop chan os.Signal) <-chan bool {
	done := make(chan bool)

	signal.Notify(stop, os.Interrupt)

	m = coffeemachine.New()
	logger = log.New(os.Stdout, "", 0)

	r := mux.NewRouter()

	s := &http.Server{
		Addr:    ":31565",
		Handler: r,
	}

	r.HandleFunc("/", machineStatus).Methods("GET")
	r.HandleFunc("/clean", clean).Methods("POST")
	r.HandleFunc("/brew", brew).Methods("POST")

	go func() {
		logger.Printf("Listening on http://localhost:31565\n")

		err := s.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	go func() {
		defer func() {
			done <- true
		}()

		<-stop

		logger.Printf("\nStopping server...\n")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := s.Shutdown(ctx)
		if err != nil {
			logger.Fatal(err)
		}

		logger.Printf("Server stopped.\nGoodbye! ☕️")
	}()

	return done
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

	err = m.Brew(coffeemachine.Beverage(c.Coffee))
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
