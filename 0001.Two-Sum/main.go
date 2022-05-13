package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 7, 11}

	fmt.Println(twoSum(nums, 5))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for k, v := range nums {
		if idx, ok := m[target - v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
