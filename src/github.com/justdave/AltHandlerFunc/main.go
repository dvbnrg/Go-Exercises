package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Messages []Message
}

func (b Book) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(b)
	}
}

func (b Book) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(b)
}

type Message struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func MakeHello(f Book) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(f)
	}
}

func main() {
	face := Book{
		Messages: []Message{
			{"Hello", "World!"},
			{"Hola", "World!"},
		},
	}
	r := mux.NewRouter()
	r.HandleFunc("/bookshelf", face.HandlerFunc)
	log.Fatal(http.ListenAndServe(":3000", r))
}
