package main

import "fmt"

func moveZeroes(nums []int)  {
	l, r, n := 0, 0, len(nums)

	for r < n {
		if nums[r] != 0 {
			nums[r], nums[l] = nums[l], nums[r]
			l++
		}
		r++
	}

}

func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}
