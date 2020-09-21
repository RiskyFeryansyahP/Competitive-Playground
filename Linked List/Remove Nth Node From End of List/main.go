package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	firstPointer := head
	secondPointer := head

	for i := 0; i < n; i++ {
		firstPointer = firstPointer.Next
	}

	if firstPointer == nil {
		return secondPointer.Next
	}

	for firstPointer.Next != nil {
		firstPointer = firstPointer.Next
		secondPointer = secondPointer.Next
	}

	secondPointer.Next = secondPointer.Next.Next
	return head
}
