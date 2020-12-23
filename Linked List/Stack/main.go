package main

import "fmt"

type Node struct {
	Value int
}

type Stack struct {
	Nodes []*Node
	size  int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(node *Node) {
	s.Nodes = append(s.Nodes, node)
	s.size++
}

func main() {
	fmt.Println("Stack Linked List")

	stack := NewStack()
	stack.Push(&Node{Value: 2})
	stack.Push(&Node{Value: 4})
	stack.Push(&Node{Value: 6})

	fmt.Println(stack.Nodes[0].Value, stack.Nodes[1].Value)
}
