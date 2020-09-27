package main

import "fmt"

// Node hold value next and prev
type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

// MyLinkedList is hold all value linked list
type MyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func main() {
	fmt.Println("Design Double Linked List!")

	node := Constructor()
	node.AddAtTail(4)

	node.display()
}

/** Initialize your data structure here. */

// Constructor is initialize new linked list
func Constructor() MyLinkedList {
	return MyLinkedList{
		Size: 0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */

// Get is get value of linked list in specific index
func (m *MyLinkedList) Get(index int) int {
	currentNode := m.Head

	if index > m.Size {
		return -1
	}

	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}

	return currentNode.Value
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */

// AddAtHead add new node in head of linked list
func (m *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{
		Value: val,
	}

	head := m.Head

	if head == nil {
		m.Head = newNode
		m.Tail = newNode
		return
	}

	head.Prev = newNode
	newNode.Next = head
	m.Head = newNode
	m.Size++
}

/** Append a node of value val to the last element of the linked list. */

// AddAtTail add new node in tail of linked list
func (m *MyLinkedList) AddAtTail(val int) {
	newNodeTail := &Node{
		Value: val,
	}

	currentNode := m.Head

	if m.Tail == nil {
		m.Head = newNodeTail
		m.Tail = newNodeTail
		return
	}

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = newNodeTail
	newNodeTail.Prev = currentNode
	m.Tail = newNodeTail
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

	// if insert at beginning
	if index == 0 {
		m.AddAtHead(val)
		return
	}

	if index == m.Size+1 {
		for i := 1; i < index; i++ {
			currentNode = currentNode.Next
		}

		currentNode.Next = newNode
		newNode.Prev = currentNode
		m.Tail = newNode
		m.Size++
		return
	}

	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}

	tmp := currentNode.Prev
	currentNode.Prev = newNode
	newNode.Next = currentNode
	newNode.Prev = tmp
	tmp.Next = newNode
	m.Size++
}

/** Delete the index-th node in the linked list, if the index is valid. */

// DeleteAtIndex delete node in specific index
func (m *MyLinkedList) DeleteAtIndex(index int) {
	currentNode := m.Head

	if index > m.Size {
		return
	}

	if index == 0 {
		m.Head = currentNode.Next
		m.Size--
		return
	}

	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}

	if index == m.Size {
		nodePrev := currentNode.Prev
		nodePrev.Next = currentNode.Next
		m.Tail = nodePrev
		m.Size--
		return
	}

	nodePrev := currentNode.Prev
	nodeNext := currentNode.Next
	nodePrev.Next = currentNode.Next
	nodeNext.Prev = nodePrev
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
