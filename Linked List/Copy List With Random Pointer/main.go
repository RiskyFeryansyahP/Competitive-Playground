package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

type LinkedList struct {
	Node *Node
}

func NewLinkedList(node *Node) *LinkedList {
	return &LinkedList{
		Node: node,
	}
}

func main() {
	dictionary := make(map[*Node]int)

	thirdNode := &Node{
		Val: 10,
	}
	secondNode := &Node{
		Val:  5,
		Next: thirdNode,
	}
	firstNode := &Node{
		Val:    1,
		Next:   secondNode,
		Random: thirdNode,
	}
	forthNode := &Node{
		Val: 11,
	}

	linkedList := NewLinkedList(firstNode)

	count := 0

	for linkedList.Node != nil {
		fmt.Println(linkedList.Node.Val)

		dictionary[linkedList.Node] = count

		count++
		linkedList.Node = linkedList.Node.Next
	}

	fmt.Println("number of node", dictionary[forthNode])

	if _, found := dictionary[forthNode]; found {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func recursifWriteDictionary(head *Node, dictionary map[*Node]*Node) *Node {
	if head == nil {
		dictionary[nil] = nil
		return nil
	}

	newNode := &Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}

	dictionary[head] = newNode

	newNode.Next = recursifWriteDictionary(head.Next, dictionary)

	return newNode
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	dictionary := make(map[*Node]*Node)

	dummyNode := recursifWriteDictionary(head, dictionary)

	root := head
	currentNode := dummyNode

	for root != nil {
		dummyNode.Random = dictionary[root.Random]

		root = root.Next
		dummyNode = dummyNode.Next
	}

	return currentNode
	/* if head == nil {
	       return nil
	   }

	   dictionary := make(map[*Node]*Node)

	   dummy := &Node{
	       Val: 0,
	   }
	   currentDummy := dummy
	   var newNode *Node

	   for head != nil {
	       if _, found := dictionary[head]; found {
	           newNode = dictionary[head]
	       } else {
	           newNode = &Node{
	               Val: head.Val,
	           }
	           dictionary[head] = newNode
	       }

	       if head.Random != nil {
	           if _, found := dictionary[head.Random]; found  {
	               newNode.Random = dictionary[head.Random]
	           } else {
	               newNode.Random = &Node{
	                   Val: head.Random.Val,
	               }
	               dictionary[head.Random] = newNode
	           }
	       }

	       currentDummy.Next = newNode
	       currentDummy = newNode

	       head = head.Next
	   }

	   return dummy.Next */

}
