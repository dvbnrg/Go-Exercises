package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Add(value1 int, value2 int) int {
	return value1 + value2
}

func Subtract(value1 int, value2 int) int {
	return value1 - value2
}

func CreateEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Item Created"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/create", CreateEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}
