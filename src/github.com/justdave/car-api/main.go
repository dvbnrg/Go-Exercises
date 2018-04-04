package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type engine struct {
	Configuration string
	EPARating     string
}

type car struct {
	engine
	Miles int
}

type lamborghini struct {
	car
	Owner string
}

type porsche struct {
	car
	Owner string
}

func fly(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am but a leaf in the wind watch me Soar!")
}

func print(w http.ResponseWriter, r *http.Request) {
}

func main() {
	gallardo := lamborghini{
		car{
			engine{
				"V10",
				"5 MPG",
			},
			1200,
		},
		"Kanye West",
	}

	carrera := porsche{
		car{
			engine{
				"Twin Turbo V6",
				"22 MPG",
			},
			14386,
		},
		"Some Doctor Guy",
	}

	r := mux.NewRouter()
	r.HandleFunc("/fly", fly).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", r))
}
