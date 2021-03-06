package main

import "fmt"

func dailyTemperatures(T []int) []int {
	length := len(T)
	ans := make([]int, length)
	stack := []int{}

	for i := 0; i < length; i++ {
		tmp := T[i]
		for len(stack) > 0 && tmp > T[stack[len(stack) - 1]] {
			prevIndex := stack[len(stack) - 1]

			stack = stack[:len(stack) - 1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}

	return ans
}

func main() {
	num := []int{73,74,75,71,69,72,76,73}
	fmt.Println(dailyTemperatures(num))
}
