package main

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	if slow == nil || fast == nil {
		return false
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == nil {
			return false
		}

		if fast == slow {
			return true
		}
	}

	return false
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/* func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	if slow == nil || fast == nil {
		return false
	}

	for slow != fast {
		if fast.Next.Next == nil || slow.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
} */
