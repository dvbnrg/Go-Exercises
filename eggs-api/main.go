package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type food struct {
	Price int    `json:"price"`
	Type  string `json:"type"`
}

type eggs struct {
	food         food `json:"food"`
	Availability bool `json:"availability"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/eggs", eggHandler).Methods("GET")
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func eggHandler(w http.ResponseWriter, r *http.Request) {
	egg := eggs{
		food:         food{Price: 1, Type: "normal"},
		Availability: true,
	}
	js, err := json.Marshal(egg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
