package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	headDummy := dummyNode

	carry := 0

	for {
		if l1 == nil && l2 == nil {
			break
		}

		x := 0
		y := 0

		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum := x + y + carry

		newNode := &ListNode{
			Val: sum % 10,
		}

		carry = sum / 10

		dummyNode.Next = newNode
		dummyNode = dummyNode.Next
	}

	if carry > 0 {
		newNode := &ListNode{
			Val: carry,
		}
		dummyNode.Next = newNode
		dummyNode = dummyNode.Next
	}

	return headDummy.Next
}
