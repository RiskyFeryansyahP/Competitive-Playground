package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	prev := &Node{
		Val: -1,
	}

	stack := []*Node{
		root,
	}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curr.Prev = prev
		prev.Next = curr
		prev = curr
		if curr.Next != nil {
			stack = append(stack, curr.Next)
		}
		if curr.Child != nil {
			stack = append(stack, curr.Child)
			curr.Child = nil
		}
	}
	root.Prev = nil
	return root
}
