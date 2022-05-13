package main

import (
	"LeetCode/structures"
	"fmt"
)

type ListNode = structures.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

func main() {
	l := getIntersectionNode(structures.Ints2List([]int{4,1,8,4,5}), structures.Ints2List([]int{5,6,1,8,4,5}))
	fmt.Println(structures.List2Ints(l))
}
