package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type nameResponse struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type jokeResponse struct {
	Type  string `json:"type"`
	Value value  `json:"value"`
}

type value struct {
	ID   int    `json:"id"`
	Joke string `json:"joke"`
}

func main() {
	port := 8080

	r := mux.NewRouter()
	r.HandleFunc("/", customJokeHandler)
	r.HandleFunc("/joke", customJokeHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Print(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}

func customJokeHandler(w http.ResponseWriter, r *http.Request) {

	name, err := http.Get("http://uinames.com/api/")

	if err != nil {
		log.Printf("name url error: %v", err)
	}

	body, err := ioutil.ReadAll(name.Body)

	n, err := getName(body)

	if err != nil {
		log.Printf("name body read error: %v", err)
	}

	jokeURL := fmt.Sprintf("http://api.icndb.com/jokes/random?firstName=%v&lastName=%v", n.Name, n.Surname)

	joke, err := http.Get(jokeURL)

	if err != nil {
		log.Printf("joke url error: %v", err)
	}

	body, err = ioutil.ReadAll(joke.Body)

	if err != nil {
		log.Printf("joke body read error: %v", err)
	}

	j, err := getJoke(body)

	if err != nil {
		log.Printf("joke body parse error: %v", err)
	}

	if j.Type == "success" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}

	defer name.Body.Close()
}

func getJoke(body []byte) (*jokeResponse, error) {
	var j = new(jokeResponse)
	err := json.Unmarshal(body, &j)
	if err != nil {
		log.Printf("joke unmarshal error: %v", err)
	}
	return j, nil
}

func getName(body []byte) (*nameResponse, error) {
	var n = new(nameResponse)
	err := json.Unmarshal(body, &n)
	if err != nil {
		log.Printf("name unmarshal error: %v", err)
	}
	return n, nil
}
