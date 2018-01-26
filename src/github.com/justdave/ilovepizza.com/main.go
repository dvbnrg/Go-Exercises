package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type customer struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Email     string `json:"email,omitempty"`
}

var customers []customer

func main() {
	router := mux.NewRouter()
	customers = append(customers, customer{Email: "asdf", Firstname: "asdf", Lastname: "asdf", Phone: "1800"})
	customers = append(customers, customer{Email: "qwer", Firstname: "qwer", Lastname: "qwer", Phone: "1900"})
	router.HandleFunc("/customer", readAll).Methods("GET")
	router.HandleFunc("/customer/{phone}", read).Methods("GET")
	router.HandleFunc("/customer/{phone}", create).Methods("PUT")
	router.HandleFunc("/customer/{phone}", delete).Methods("DELETE")
	router.HandleFunc("/export", dumpcsv).Methods("GET")
	router.HandleFunc("/import", grabcsv).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func create(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var customer customer
	_ = json.NewDecoder(req.Body).Decode(&customer)
	customer.Phone = params["phone"]
	customers = append(customers, customer)
	json.NewEncoder(w).Encode(customers)
}

func read(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range customers {
		if item.Phone == params["phone"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&customer{})
}

func readAll(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(customers)
}

func delete(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range customers {
		if item.Phone == params["phone"] {
			customers = append(customers[:index], customers[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(customers)
}

func dumpcsv(w http.ResponseWriter, req *http.Request) {
	file, err := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()

	if err != nil {
		os.Exit(1)
	}
	x := []string{"Country", "City", "Population"}
	y := []string{"Japan", "Tokyo", "923456"}
	z := []string{"Australia", "Sydney", "789650"}
	csvWriter := csv.NewWriter(file)
	strWrite := [][]string{x, y, z}
	csvWriter.WriteAll(strWrite)
	csvWriter.Flush()
}

func grabcsv(w http.ResponseWriter, req *http.Request) {

}
