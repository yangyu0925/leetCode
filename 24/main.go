package main

import (
	"LeetCode/structures"
	"fmt"
)

func swapPairs(head *structures.ListNode) *structures.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head

	return newHead
}

func main()  {
	head := swapPairs(structures.Ints2List([]int{1, 2, 3 ,4}))

	fmt.Println(structures.List2Ints(head))
}
