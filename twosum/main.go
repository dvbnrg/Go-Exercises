package main

import "fmt"

func main() {

	sampleArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	x := twoSum(sampleArray, 8)

	fmt.Println(x)
}

func twoSum(nums []int, target int) []int {
	// map from int -> int
	m := make(map[int]int)

	// python like enumerate
	for i, num := range nums {
		// get value and check existence in map
		// and in if statement where ok is the boolean comparison value
		if index, ok := m[target-num]; ok {
			// construct array and return values
			// Go also has multiple returns, but LeetCode requested this format to stay consistent with other languages.
			return []int{index, i}
		}
		// python like set in map
		m[num] = i
	}
	return nil
}
