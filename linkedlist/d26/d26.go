package main

import "fmt"

func main() {
	input := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	ans := reverseList(&input)
	fmt.Println(ans)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// set prev to nil
	var prev *ListNode
	// set curr to head
	current := head

	// while curr is not nil
	for current != nil {
		// set nextTemp to curr.Next
		nextTemp := current.Next
		// set current.Next to prev
		current.Next = prev
		// set prev to curr
		prev = current
		// set curr to nextTemp
		current = nextTemp
	}
	return prev
}
