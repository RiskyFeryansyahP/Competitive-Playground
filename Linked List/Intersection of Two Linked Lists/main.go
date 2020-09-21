package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pointerA := headA
	pointerB := headB

	if pointerA == nil || pointerB == nil {
		return nil
	}

	// check if in the start pointerA and pointerB placed in the same node
	if pointerA == pointerB {
		return pointerA
	}

	for pointerA != pointerB {
		pointerA = pointerA.Next
		pointerB = pointerB.Next

		if pointerA == nil {
			pointerA = headB
		}

		if pointerB == nil {
			pointerB = headA
		}

		if pointerA == pointerB {
			return pointerA
		}

		if pointerA.Next == nil && pointerB.Next == nil {
			break
		}

	}

	return nil
}

// this optimal solution
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}

		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}
