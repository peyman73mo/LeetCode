package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	tail := head

	for {
		val1, val2 := 0, 0

		if l1.Next != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2.Next != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		num := val1 + val2 + tail.Val
		extra := num / 10
		tail.Val = num % 10

		if l1 == nil && l2 == nil && extra == 0 {
			break
		}

		tail.Next = &ListNode{
			Val:  extra,
			Next: nil,
		}
		tail = tail.Next

	}
	return head
}
func main() {

}
