package main

import (
	"fmt"
	"math"
)


type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	sum := 0
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {

		var1, var2 := 0, 0
		if l1 != nil {
			var1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			var2 = l2.Val
			l2 = l2.Next
		}

		sum = var1 + var2 + carry
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return head.Next
}

func creat_number(l *ListNode) int {
	var num int = 0
	multiple := 0
	for l != nil {
		num = num + (l.Val)*int(math.Pow10(multiple))
		l = l.Next
		multiple++
	}
	return num
}

func create_linklist(num int) *ListNode {
	head := &ListNode{}
	current := head
	for num != 0 {
		current.Val = num % 10
		num = num / 10
		if num != 0 {
			current.Next = &ListNode{}
			current = current.Next
		}
	}
	return head
}

func print_result(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	// fmt.Println(addTwoNumbers(l1, l2))
	print_result(addTwoNumbers(l1, l2))
	l1 = &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}
	l2 = &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}}
	print_result(addTwoNumbers(l1, l2))
	// [9,9,9,9,9,9,9]
	// [9,9,9,9]
	l1 = &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}}}
	l2 = &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}
	print_result(addTwoNumbers(l1, l2))

}
