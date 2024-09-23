package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var new *ListNode // new is nil
	current := head
	next := current

	for {
		if current == nil {
			break
		}
		next = next.Next
		current.Next = new

		new = current
		current = next
	}

	return new
}
