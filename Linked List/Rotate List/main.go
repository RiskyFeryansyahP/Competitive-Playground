package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	size := 1

	tail := head

	for tail.Next != nil {
		tail = tail.Next
		size++
	}

	k = k % size
	rotateStep := size - k

	if k == 0 {
		return head
	}

	tail.Next = head

	for rotateStep > 0 {
		tail = tail.Next
		rotateStep = rotateStep - 1
	}

	head = tail.Next
	tail.Next = nil

	return head
}
