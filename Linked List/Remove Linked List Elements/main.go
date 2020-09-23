package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummyList := &ListNode{-1, head}
	prev := dummyList

	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
			continue
		}
		prev = prev.Next
	}

	return dummyList.Next
}
