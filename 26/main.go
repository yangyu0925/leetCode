package main

import "fmt"

func remobeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0

	for last < len(nums) - 1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last + 1] = nums[finder]
		last++
	}
	return last + 1
}

func main() {
	fmt.Println(remobeDuplicates([]int{0,0,1,1,1,2,2,3,3,4,5,6,7}))
}
