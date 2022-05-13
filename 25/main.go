package main

import (
	"LeetCode/structures"
	"fmt"
)

type ListNode = structures.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head

	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}

	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(first, last *ListNode) *ListNode {
	prev := last

	for first != last {
		tmp := first.Next
		first.Next = prev
		prev = first
		first = tmp
	}

	return prev
}

func main()  {
	l := reverseKGroup(structures.Ints2List([]int{1, 2, 3, 4, 5}), 2);
	fmt.Println(structures.List2Ints(l))
}
