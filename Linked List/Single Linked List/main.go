package main

import "fmt"

// Node is hold all Node needed
type Node struct {
	Value int
	Next  *Node
}

// LinkedList is create a linked list
type LinkedList struct {
	Head *Node
}

func main() {
	fmt.Println("Single Linked List")

	linkedList := NewLinkedList(10)
	linkedList.addNode(20)
	linkedList.addNode(30)
	linkedList.addNode(40)
	linkedList.addNode(50)

	linkedList.display()
	fmt.Println("----------------")

	// res := linkedList.searchLinkedList(60)
	// fmt.Println("Linked List finded in number", res)

	// linkedList.addNodeSpecific(20, 70)
	// fmt.Println("Spesific", linkedList.Head.Next.Next.Next.Value)

	linkedList.deleteNode(10)
	linkedList.display()
}

// NewLinkedList is create a new node linked list
func NewLinkedList(value int) *LinkedList {
	return &LinkedList{
		Head: &Node{
			Value: value,
			Next:  nil,
		},
	}
}

func (l *LinkedList) addNode(val int) {
	newNode := &Node{
		Value: val,
	}

	// check if head is nil, fill the head first
	if l.Head == nil {
		l.Head = newNode
	}

	node := l.Head

	for node.Next != nil {
		node = node.Next
	}

	node.Next = newNode
}

func (l *LinkedList) deleteNode(value int) {
	head := l.Head

	if head.Value == value {
		l.Head = head.Next
		return
	}

	for head != nil {
		if head.Next.Value == value {
			break
		}

		head = head.Next
	}

	head.Next = head.Next.Next
}

func (l *LinkedList) display() {
	head := l.Head

	no := 1

	for head != nil {
		fmt.Printf("Value %d: %d \n", no, head.Value)

		head = head.Next

		no++
	}
}

func (l *LinkedList) searchLinkedList(value int) int {
	currentLinkedList := l.Head

	no := 1

	for currentLinkedList != nil {
		if currentLinkedList.Value == value {
			return no
		}
		currentLinkedList = currentLinkedList.Next
		no++
	}
	return 0
}

func (l *LinkedList) addNodeSpecific(value, newValue int) {
	newNode := &Node{
		Value: newValue,
	}

	head := l.Head

	for head != nil {
		if head.Value == value {
			break
		}

		head = head.Next
	}

	currentNode := head.Next
	newNode.Next = currentNode
	head.Next = newNode
}
