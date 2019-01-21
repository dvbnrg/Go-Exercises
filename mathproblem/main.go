package main

import "fmt"

func main() {
	fmt.Println(isMultiple(1000))

	fmt.Println("-----------")

	fmt.Println(sumAll(isMultiple(1000)))
}

func isMultiple(initial int) (multiples []int) {

	num := make([]int, initial+1)

	for x := range num {
		num[x] = x
		x++
	}

	for _, x := range num {
		if x%3 == 0 {
			multiples = append(multiples, x)
		} else if x%5 == 0 {
			multiples = append(multiples, x)
		}
	}
	return
}

func sumAll(num []int) (result int) {
	for _, x := range num {
		result += x
	}
	return
}
