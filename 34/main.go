package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 6, 7, 8, 8, 9 ,10}

	fmt.Println(searchRange(nums, 8))
}

func searchRange(nums []int, target int) []int {
	n := len(nums)
	left := 0
	right := n - 1

	for left <= right {
		m := (left + right) / 2
		if nums[m] == target {
			i, j := m, m
			for (i >= 0 && nums[i] == target) || (j <= n - 1 && nums[j] == target) {
				if i >= 0 && nums[i] == target {
					i--
				}
				if j <= n - 1 && nums[j] == target {
					j++
				}
			}
			return []int{i + 1, j - 1}
		}
		if nums[m] > target {
			right = m - 1
		} else {
			left = m + 1
		}
	}

	return []int{-1, -1}
}
