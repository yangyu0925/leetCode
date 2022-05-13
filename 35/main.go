package main

import "fmt"

func searchInsert(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n - 1
	ans := n
	for l <= r {
		mid := (l + r) / 2
		if target <= nums[mid] {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(searchInsert([]int{1, 2, 4, 5, 6,}, 3))
}
