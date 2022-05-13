package main

import "LeetCode/structures"

func hasCycle(head *structures.ListNode) bool {
	if head == nil {
		return false
	}

	m := make(map[*structures.ListNode]struct{})

	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}

	return false
}


func hasCycle1(head *structures.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for slow != fast{

		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
