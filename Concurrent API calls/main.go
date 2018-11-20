package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Response struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
}

var wg sync.WaitGroup

func main() {

	url := "https://api.github.com/users/octocat"
	url2 := "https://api.github.com/users/dvbnrg"

	wg.Add(2)
	go fetch(url)
	go fetch(url2)
	wg.Wait()
}

func fetch(url string) {
	defer wg.Done()
	res, err := http.Get(url)

	if err != nil {
		log.Printf("name url error: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Printf("name body read error: %v", err)
	}

	n := new(Response)

	err = json.Unmarshal(body, &n)

	if err != nil {
		log.Print(err)
	}

	fmt.Println(n)
}
