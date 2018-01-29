package main

import (
	"encoding/json"
	"fmt"
)

type customer struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Email     string `json:"email,omitempty"`
}

var customers []customer

func main() {
	customers = append(customers, customer{Email: "asdf", Firstname: "asdf", Lastname: "asdf", Phone: "1800"})
	customers = append(customers, customer{Email: "qwer", Firstname: "qwer", Lastname: "qwer", Phone: "1900"})
	fmt.Println("Hello")
	fmt.Println(customers)
	parseobject(customers)
}

func tostring(c customer) string {
	out, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func parseobject(c []customer) [][]string {
	result := [][]string{}
	for i := 0; i < len(c); i++ {
		fmt.Println(c[i].Email + c[i].Firstname + c[i].Lastname + c[i].Phone)
	}
	return result
}
