package main

import "fmt"

func main() {
	n := 2
	input := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	ans := removeNthFromEnd(&input, n)
	fmt.Println(ans)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	a, b := head, head

	for n > 0 {
		b = b.Next
		n--
	}

	if b == nil {
		return a.Next
	}

	for b.Next != nil {
		a = a.Next
		b = b.Next
	}
	a.Next = a.Next.Next
	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  func removeNthFromEnd(head *ListNode, n int) *ListNode {
//     if head == nil {
//         return head
//     }

//     slow, fast := head, head
//     for n > 0 {
//         fast = fast.Next
//         n--
//     }

//     if fast == nil {
//         return slow.Next
//     }

//     for fast.Next != nil {
//         slow = slow.Next
//         fast = fast.Next
//     }

//     slow.Next = slow.Next.Next
//     return head
// }
