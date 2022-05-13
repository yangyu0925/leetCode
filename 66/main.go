package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits) - 1

	for i := n; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)

}

func main() {
	digits := []int{1, 2, 3, 9}

	fmt.Println(plusOne(digits))

}
