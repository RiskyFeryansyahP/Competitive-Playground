package main

import "fmt"

// DoubleLinkedList is hold data for linked list
type DoubleLinkedList struct {
	Val  int
	Next *DoubleLinkedList
	Prev *DoubleLinkedList
}

func main() {
	fmt.Println("Hello World!")

	node := &DoubleLinkedList{
		Val: 7,
	}

	node1 := addDoubleLinkedList(node, 8)
	_ = addDoubleLinkedList(node1, 10)

	fmt.Println(node.Next.Next)

	_ = addSpecificPosition(node, 8, 9)

	fmt.Println(node.Next.Next.Next.Prev)

}

func addDoubleLinkedList(node *DoubleLinkedList, val int) *DoubleLinkedList {
	newNode := &DoubleLinkedList{
		Val: val,
	}

	head := node

	if head == nil {
		head = newNode
		return head
	}

	for head.Next != nil {
		head = head.Next
	}

	head.Next = newNode
	newNode.Prev = head

	return head
}

func addSpecificPosition(node *DoubleLinkedList, findVal int, val int) *DoubleLinkedList {
	newNode := &DoubleLinkedList{
		Val: val,
	}

	head := node

	if head == nil {
		head = newNode
		return head
	}

	for head != nil {
		if head.Val == findVal {
			break
		}
		head = head.Next
	}

	tmp := head.Next
	head.Next = newNode
	newNode.Prev = head
	newNode.Next = tmp
	tmp.Prev = newNode

	return head
}
