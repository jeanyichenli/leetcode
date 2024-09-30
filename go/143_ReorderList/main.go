package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	midNode := findMiddleNode(head)

	newhead := midNode.Next
	midNode.Next = nil

	// reverse newhead
	newhead = reverseList(newhead)

	// reorder
	var tmp1, tmp2 *ListNode

	for head != nil && newhead != nil {
		tmp1 = head.Next
		tmp2 = newhead.Next

		head.Next = newhead
		newhead.Next = tmp1

		head = tmp1
		newhead = tmp2
	}

}

func findMiddleNode(head *ListNode) *ListNode {
	// two pointers: slow & fast
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseList(head *ListNode) *ListNode {
	var newhead *ListNode

	cur, next := head, head

	for cur != nil {
		next = cur.Next

		cur.Next = newhead
		newhead = cur
		cur = next
	}

	return newhead
}
