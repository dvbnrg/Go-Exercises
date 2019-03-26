package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Request struct {
	Data []Data `json:"data"`
}

type Data struct {
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	Price          string `json:"price"`
	YesterdayClose string `json:"close_yesterday"`
	Currency       string `json:"currency"`
	MarketCap      string `json:"market_cap"`
	Volume         string `json:"volume"`
	Timezone       string `json:"timezone"`
	TimezoneName   string `json:"timezone_name"`
	GmtOffset      string `json:"gmt_offset"`
	LastTradeTime  string `json:"last_trade_time"`
	Market         string `json:"stock_exchange_short"`
}

const apiKey = ""

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/stock", defaultsearch).Methods("GET")
	router.HandleFunc("/stock/{symbol}", search).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func defaultsearch(w http.ResponseWriter, req *http.Request) {
	// symbol := mux.Vars(req)

	url := "https://www.worldtradingdata.com/api/v1/stock?symbol=AXP&api_token="
	url += apiKey

	x := fetch(url)
	json.NewEncoder(w).Encode(&x)
}

func search(w http.ResponseWriter, req *http.Request) {
	symbol := mux.Vars(req)

	url := "https://www.worldtradingdata.com/api/v1/stock?symbol="
	url += symbol["symbol"]
	url += "&api_token="
	url += apiKey

	x := fetch(url)
	json.NewEncoder(w).Encode(&x)
}

func fetch(url string) (n *Request) {
	res, err := http.Get(url)

	if err != nil {
		log.Printf("name url error: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Printf("name body read error: %v", err)
	}

	n = new(Request)

	err = json.Unmarshal(body, &n)

	if err != nil {
		log.Printf("unmarshal error: %v", err)
	}

	return
}
