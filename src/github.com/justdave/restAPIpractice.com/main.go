package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var m map[string]string

func main() {
	m = make(map[string]string)
	m["key"] = "value"
	//fmt.Println(m["key"])

	//var h http.Handler = http.HandlerFunc(home)
	r := mux.NewRouter()
	// This will assign the home page to the
	// NotFoundHandler
	//r.NotFoundHandler = h
	r.HandleFunc("/{Key}", get).Methods("GET")
	r.HandleFunc("/{Key}", put).Methods("PUT")
	http.ListenAndServe(":3000", r)
}

func get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Key"]
	if key != "" {
		value, ok := m[key]
		if !ok {
			w.WriteHeader(404)
			return
		}
		w.Write([]byte(value))
		//fmt.Println("value: ", key)
	} else {
		fmt.Println("404 key not found")
	}
}

func put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Key"]
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	m[key] = string(b)
}
