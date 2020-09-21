package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head

	if slow == nil || fast == nil {
		return nil
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// check if fast nil, there no have cycle
		if fast == nil {
			return nil
		}

		// check if fast == slow, then there have a cycle
		if fast == slow {
			slow = head
			break
		}
	}

	for fast.Next != nil {
		if fast == slow {
			return fast
		}

		fast = fast.Next
		slow = slow.Next
	}

	return nil
}
