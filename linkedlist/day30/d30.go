package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	println(divmod(12, 10))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy *ListNode
	cur := dummy
	var carry int

	for (l1.Next != nil) || (l2.Next != nil) || (carry != 0) {
		var v1, v2 int
		v1 = l1.Val
		if l1 != nil {
			v1 = 0
		}
		v2 = l2.Val
		if l2 != nil {
			v2 = 0
		}
		val := v1 + v2 + carry
		a, b := divmod(val, 10)
		carry = b
		cur.Next = &ListNode{a, &ListNode{}}
	}
	return dummy.Next
}

func divmod(num, denom int) (quotient, remainder int) {
	quotient = num / denom
	remainder = num % denom
	return
}
