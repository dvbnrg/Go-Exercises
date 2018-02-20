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

type Transform [3][3]float64 // A 3x3 array, really an array of arrays.
type LinesOfText [][]string  // A slice of string slices.

var customers []customer

func main() {
	customers = append(customers, customer{Email: "asdf", Firstname: "as", Lastname: "df", Phone: "1800"})
	customers = append(customers, customer{Email: "qwer", Firstname: "qw", Lastname: "er", Phone: "1900"})
	customers = append(customers, customer{Email: "zxcv", Firstname: "zc", Lastname: "cv", Phone: "2000"})
	fmt.Println("Hello")
	fmt.Println(customers)
	parseobject(customers)

	//2d Array Practice
	text := LinesOfText{
		[]string{"Now is the time"},
		[]string{"for all good gophers"},
		[]string{"to bring some fun to the party."},
	}

	fmt.Println(text)

	YSize := 6
	XSize := 4
	// Allocate the top-level slice.
	picture := make([][]uint8, YSize) // One row per unit of y.
	// Loop over the rows, allocating the slice for each row.
	for i := range picture {
		picture[i] = make([]uint8, XSize)
	}

	fmt.Println(picture)

	// Allocate the top-level slice, the same as before.
	picture2 := make([][]uint8, YSize) // One row per unit of y.
	// Allocate one large slice to hold all the pixels.
	pixels := make([]uint8, XSize*YSize) // Has type []uint8 even though picture is [][]uint8.
	// Loop over the rows, slicing each row from the front of the remaining pixels slice.
	for i := range picture {
		picture2[i], pixels = pixels[:XSize], pixels[XSize:]
	}

	fmt.Println(pixels)
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
