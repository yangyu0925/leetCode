package main

import "fmt"

func maxProfit(nums []int) int {
	length := len(nums)
	if length < 2 {
		return 0
	}

	min, ans := nums[0], 0

	for i := 1; i < length; i++ {
		if nums[i] - min > ans {
			ans = nums[i] - min
		}

		if nums[i] < min {
			min = nums[i]
		}
	}

	return ans
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 2, 3, 6, 4}))
}
