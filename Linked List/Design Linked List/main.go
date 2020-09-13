package main

import (
	"fmt"
)

func main() {
	fmt.Println("Design Linked List")

	linkedList := Constructor()
	fmt.Println("My Linked List", linkedList)

	linkedList.AddAtHead(9)
	linkedList.Get(1)
	linkedList.AddAtIndex(1, 1)
	linkedList.AddAtIndex(1, 7)
	linkedList.DeleteAtIndex(1)
	linkedList.AddAtHead(7)
	linkedList.AddAtHead(4)
	linkedList.DeleteAtIndex(1)
	linkedList.AddAtIndex(1, 4)
	linkedList.AddAtHead(2)
	linkedList.DeleteAtIndex(5)

	linkedList.display()

	fmt.Println("Size of linked list", linkedList.Size)
}

// Node is reference of node linked list
type Node struct {
	Value int
	Next  *Node
}

// MyLinkedList is hold all value linked list
type MyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

/** Initialize your data structure here. */

// Constructor is initialize new linked list
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */

// Get is get value of linked list in specific index
func (m *MyLinkedList) Get(index int) int {
	currentNode := m.Head

	nodeIndex := 0

	if index > m.Size {
		return -1
	}

	for currentNode != nil {
		if nodeIndex == index {
			break
		}

		currentNode = currentNode.Next
		nodeIndex++
	}

	return currentNode.Value
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */

// AddAtHead add new node in head of linked list
func (m *MyLinkedList) AddAtHead(val int) {
	node := &Node{
		Value: val,
	}

	if m.Head == nil {
		m.Head = node
		return
	}

	currentHead := m.Head
	m.Head = node
	m.Head.Next = currentHead
	m.Size++
}

/** Append a node of value val to the last element of the linked list. */

// AddAtTail add new node in tail of linked list
func (m *MyLinkedList) AddAtTail(val int) {
	newNodeTail := &Node{
		Value: val,
	}

	currentNode := m.Head

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = newNodeTail
	m.Tail = currentNode.Next
	m.Size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */

// AddAtIndex add new node in specific index
func (m *MyLinkedList) AddAtIndex(index int, val int) {
	newNode := &Node{
		Value: val,
	}

	currentNode := m.Head

	if index > m.Size+1 {
		return
	}

	nodeIndex := 1

	if index == 0 {
		currentHeadNode := m.Head
		m.Head = newNode
		m.Head.Next = currentHeadNode

		return
	}

	for currentNode != nil {
		if nodeIndex == index {
			break
		}

		currentNode = currentNode.Next
		nodeIndex++
	}

	// check if index equals to the length of linked list
	if currentNode.Next == nil {
		currentNode.Next = newNode
		m.Tail = newNode
		m.Size++
		return
	}

	currentNextNode := currentNode.Next
	currentNode.Next = newNode
	currentNode.Next.Next = currentNextNode
	m.Size++
}

/** Delete the index-th node in the linked list, if the index is valid. */

// DeleteAtIndex delete node in specific index
func (m *MyLinkedList) DeleteAtIndex(index int) {
	nodeIndex := 1

	currentNode := m.Head

	if index > m.Size {
		return
	}

	if index == 0 {
		m.Head = m.Head.Next
		return
	}

	for currentNode != nil {
		if nodeIndex == index {
			break
		}

		currentNode = currentNode.Next
		nodeIndex++
	}

	currentNode.Next = currentNode.Next.Next
	m.Size--
}

func (m *MyLinkedList) display() {
	currentNode := m.Head

	fmt.Println("-----------------------")

	for currentNode != nil {
		fmt.Println(currentNode.Value)

		currentNode = currentNode.Next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
