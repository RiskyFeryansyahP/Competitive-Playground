package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head

	// 1->2->2->1
	for fast != nil {
		if fast.Next == nil {
			break
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	// 1->2->2->nil, because the slow doing reverse the linked list
	fast = head
	slow = reverseLinkedList(slow)

	// slow after reverse linkedlist 1->2->nil (simplify: 1->2->2->nil<-2<-1)
	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}

		slow = slow.Next
		fast = fast.Next
	}

	return true
}

// 1->2->3->4->nil
// nil<-1<-2<-3<-4
func reverseLinkedList(node *ListNode) *ListNode {
	currentNode := node
	var prev *ListNode
	var next *ListNode

	for currentNode != nil {
		next = currentNode.Next
		currentNode.Next = prev
		prev = currentNode
		currentNode = next
	}

	return prev
}
