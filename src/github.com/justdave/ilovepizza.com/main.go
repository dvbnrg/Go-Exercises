package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
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
	customers = append(customers, customer{Email: "zxcv", Firstname: "zc", Lastname: "cv", Phone: "2000"})
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
	file, err := os.OpenFile("dump.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()

	if err != nil {
		os.Exit(1)
	}
	csvWriter := csv.NewWriter(file)
	strWrite := parseobject(customers)
	csvWriter.WriteAll(strWrite)
	csvWriter.Flush()
}

func grabcsv(w http.ResponseWriter, req *http.Request) {

}

func tostring(c customer) string {
	out, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func parseobject(c []customer) [][]string {
	result := make([][]string, 4)
	for i := range c {
		out := c[i].Email + ", " + c[i].Firstname + ", " + c[i].Lastname + ", " + c[i].Phone
		fmt.Println("Full String:" + out)

		result[i] = make([]string, 4)
		fmt.Println("Output i: ", i)

		result[i] = append(result[i], c[i].Email)
		fmt.Println(result[i])
		result[i] = append(result[i], c[i].Firstname)
		fmt.Println(result[i])
		result[i] = append(result[i], c[i].Lastname)
		fmt.Println(result[i])
		result[i] = append(result[i], c[i].Phone)
		fmt.Println(result[i])

	}
	return result
}
