package main

import "fmt"

type ListNode struct {
	V int
	Next *ListNode
}

func getLength(head *ListNode) (length int) {
	for head != nil {
		length++
		head = head.Next
	}
	return 
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length - n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func Int2List(nums []int) *ListNode {
	head := &ListNode{}

	t := head

	for _, v := range nums{
		t.Next = &ListNode{V:v}
		t = t.Next
	}

	return head.Next
}

func List2Int(head *ListNode) []int {
	nums := make([]int, 0)

	for head != nil {
		nums = append(nums, head.V)
		head = head.Next
	}

	return nums
}

func main() {
	l := removeNthFromEnd(Int2List([]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21}), 7)
	fmt.Println(List2Int(l))
}
