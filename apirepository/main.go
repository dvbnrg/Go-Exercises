package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type repositories struct {
	Name  string `json:"full_name"`
	Owner owner  `json:"owner"`
}

type owner struct {
	URL string `json:"url"`
}

func main() {
	fmt.Println("Hello")

	name, err := http.Get("https://api.github.com/repositories")

	if err != nil {
		log.Printf("name url error: %v", err)
	}

	body, err := ioutil.ReadAll(name.Body)

	n, err := getInfo(body)

	if err != nil {
		log.Printf("name body read error: %v", err)
	}

	fmt.Println(n)
}

func getInfo(body []byte) ([]repositories, error) {
	n := make([]repositories, 0)
	err := json.Unmarshal(body, &n)
	if err != nil {
		log.Printf("name unmarshal error: %v", err)
	}
	return n, nil
}
