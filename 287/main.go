package main

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]

	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]

	}

	walker := 0
	for walker != slow {
		walker = nums[walker]
		slow = nums[slow]
	}

	return walker
}

func main() {
	findDuplicate([]int{1, 3, 4, 2, 2})
}