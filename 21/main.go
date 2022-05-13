package main

import (
	"LeetCode/structures"
	"fmt"
)

type ListNode = structures.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

func main() {
	l := mergeTwoLists(structures.Ints2List([]int{1,2,4, 5}), structures.Ints2List([]int{1,3,4}))

	fmt.Println(structures.List2Ints(l))

}
