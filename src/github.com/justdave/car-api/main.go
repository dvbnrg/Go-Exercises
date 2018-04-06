package main

import (
	"fmt"
	"net/http"
)

//Garage is the Scuderia for the Cars
type Garage struct {
	Capacity int
	Cars     []Car
}

// Engine is the basic part of the Car
type Engine struct {
	Configuration string
	EPARating     string
	Mileage       int
}

//Car is the abstract part of the Car
type Car interface {
	getMileage() int
	getEngine() Engine
	Start()
	Drive()
	Stop()
	Maintain()
}

// car is the abstract part of the Car
type car struct {
	Engine Engine
}

func (c car) getMileage() int {
	return c.Engine.Mileage
}

func (c car) getEngine() Engine {
	return c.Engine
}

func (c car) Start() {
	panic("implement me")
}

func (c car) Drive() {
	panic("implement me")
}

func (c car) Stop() {
	panic("implement me")
}

func (c car) Maintain() {
	panic("implement me")
}

//Ferrari is the more complete version
type Ferrari struct {
	car
	Owner string
	Model string
}

//Lamborghini is another version of Car
type Lamborghini struct {
	car
	Owner string
	Model string
}

func fly(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am but a leaf in the wind watch me Soar!")
}

func printCar(w http.ResponseWriter, r *http.Request) {
}

func main() {
	jaylenogarage := Garage{
		Capacity: 100,
		Cars: []Car{
			&Lamborghini{
				car: car{
					Engine: Engine{
						Mileage: 1200,
					},
				},
				Owner: "ME",
			},
			&Ferrari{
				car: car{
					Engine: Engine{
						Mileage: 1200,
					},
				},
				Owner: "DAVE",
			},
		},
	}

	for _, c := range jaylenogarage.Cars {
		fmt.Printf("%#v\n", c)
	}

	//r := mux.NewRouter()
	//r.HandleFunc("/fly", fly).Methods("GET")
	//log.Fatal(http.ListenAndServe(":3000", r))
}
