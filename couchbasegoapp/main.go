package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/couchbase/gocb"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

type Person struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
}

type N1qlPerson struct {
	Person Person `json:"person"`
}

var bucket *gocb.Bucket

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var n1qlParams []interface{}
	query := gocb.NewN1qlQuery("SELECT * FROM `restful-sample` AS person WHERE META(person).id = $1")
	params := mux.Vars(req)
	n1qlParams = append(n1qlParams, params["id"])
	rows, _ := bucket.ExecuteN1qlQuery(query, n1qlParams)
	var row N1qlPerson
	rows.One(&row)
	json.NewEncoder(w).Encode(row.Person)
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	var person []Person
	query := gocb.NewN1qlQuery("SELECT * FROM `restful-sample` AS person")
	rows, _ := bucket.ExecuteN1qlQuery(query, nil)
	var row N1qlPerson
	for rows.Next(&row) {
		person = append(person, row.Person)
	}
	json.NewEncoder(w).Encode(person)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var person Person
	var n1qlParams []interface{}
	_ = json.NewDecoder(req.Body).Decode(&person)
	query := gocb.NewN1qlQuery("INSERT INTO `restful-sample` (KEY, VALUE) VALUES ($1, {'firstname': $2, 'lastname': $3, 'email': $4})")
	n1qlParams = append(n1qlParams, uuid.NewV4().String())
	n1qlParams = append(n1qlParams, person.Firstname)
	n1qlParams = append(n1qlParams, person.Lastname)
	n1qlParams = append(n1qlParams, person.Email)
	_, err := bucket.ExecuteN1qlQuery(query, n1qlParams)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(person)
}

func UpdatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var person Person
	var n1qlParams []interface{}
	_ = json.NewDecoder(req.Body).Decode(&person)
	query := gocb.NewN1qlQuery("UPDATE `restful-sample` USE KEYS $1 SET firstname = $2, lastname = $3, email = $4")
	params := mux.Vars(req)
	n1qlParams = append(n1qlParams, params["id"])
	n1qlParams = append(n1qlParams, person.Firstname)
	n1qlParams = append(n1qlParams, person.Lastname)
	n1qlParams = append(n1qlParams, person.Email)
	_, err := bucket.ExecuteN1qlQuery(query, n1qlParams)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(person)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var n1qlParams []interface{}
	query := gocb.NewN1qlQuery("DELETE FROM `restful-sample` AS person WHERE META(person).id = $1")
	params := mux.Vars(req)
	n1qlParams = append(n1qlParams, params["id"])
	_, err := bucket.ExecuteN1qlQuery(query, n1qlParams)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&Person{})
}

func main() {
	router := mux.NewRouter()
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucket, _ = cluster.OpenBucket("restful-sample", "")
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people", CreatePersonEndpoint).Methods("PUT")
	router.HandleFunc("/people/{id}", UpdatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}
