package main

import (
	"fmt"
)

func main() {
	power := 1000
	fmt.Printf("It's over %d\n", power)

	name, power := "Goku", getPower()
	fmt.Printf("%s's power is over %d\n", name, power)
}

func getPower() int {
	return 9001
}
