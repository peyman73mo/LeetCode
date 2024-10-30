package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// create 2 sperate linked list, one for odd, one for even; then merge them
	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	// merge
	odd.Next = evenHead
	return head

}