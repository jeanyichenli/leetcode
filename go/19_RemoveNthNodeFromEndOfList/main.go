package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{Val: 0, Next: head}

	front := result
	for i := 0; i <= n; i++ {
		front = front.Next
	}

	current := result
	for front != nil {
		current = current.Next
		front = front.Next
	}

	current.Next = current.Next.Next
	return result.Next
}
