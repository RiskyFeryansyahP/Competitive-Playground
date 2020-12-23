package main

import "fmt"

// Node ...
type Node struct {
	Next  *Node
	Prev  *Node
	Key   int
	Value int
}

// LinkedList ...
type LinkedList struct {
	Head *Node
	Tail *Node
	Node *Node
}

// LRUCache ...
type LRUCache struct {
	Linkedlist *LinkedList
	Map        map[int]*Node
	Capacity   int
}

// Put ...
func (cache *LRUCache) Put(key, val int) {
	head := cache.Linkedlist.Head

	node, found := cache.Map[key]
	if found {
		moveNodeToFront(node, cache)
		node.Value = val
		return
	}

	if cache.Linkedlist.Head == nil || cache.Linkedlist.Tail == nil {
		newNode := &Node{
			Value: val,
			Next:  nil,
			Prev:  nil,
		}
		dummyTail := &Node{
			Next: nil,
			Prev: nil,
		}

		cache.Linkedlist.Head = newNode
		cache.Linkedlist.Tail = newNode
		cache.Linkedlist.Tail.Next = dummyTail
		cache.Map[key] = newNode
		return
	}

	if len(cache.Map) >= cache.Capacity {
		// doing evict least recently used
		evict(cache)
		fmt.Println("delete key", key)
	}

	newNode := &Node{
		Value: val,
		Next:  head,
		Key:   key,
		Prev:  nil,
	}

	cache.Linkedlist.Head.Prev = newNode
	cache.Linkedlist.Head = newNode
	cache.Map[key] = newNode
}

// Get ...
func (cache *LRUCache) Get(key int) int {
	node, found := cache.Map[key]
	if found {
		node := moveNodeToFront(node, cache)

		return node.Value
	}

	return -1
}

func moveNodeToFront(node *Node, mark *LRUCache) *Node {
	if node == mark.Linkedlist.Head {
		return node
	}

	if node == mark.Linkedlist.Tail {
		mark.Linkedlist.Tail = node.Prev
	}

	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev

	mark.Linkedlist.Head.Prev = node
	node.Next = mark.Linkedlist.Head
	mark.Linkedlist.Head = node

	return node
}

func evict(cache *LRUCache) {
	delete(cache.Map, cache.Linkedlist.Tail.Key)
	tail := cache.Linkedlist.Tail
	next := cache.Linkedlist.Tail.Next
	cache.Linkedlist.Tail = tail.Prev
	cache.Linkedlist.Tail.Next = next
}

func displayLinkedList(cache *LRUCache) {
	currentNode := cache.Linkedlist.Head
	for currentNode.Next != nil {
		fmt.Print(currentNode.Value, " -> ")
		currentNode = currentNode.Next
	}

	fmt.Print("nil \n")
}

func main() {
	fmt.Println("LRU Implementation")

	linkedList := &LinkedList{
		Head: nil,
		Node: nil,
		Tail: nil,
	}

	cache := &LRUCache{
		Linkedlist: linkedList,
		Map:        make(map[int]*Node),
		Capacity:   5,
	}

	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(3, 5)
	cache.Put(4, 7)
	cache.Put(5, 9)
	cache.Get(1)
	cache.Get(1)
	cache.Get(5)
	cache.Put(6, 11)
	cache.Get(3)
	cache.Put(3, 10)

	displayLinkedList(cache)
}
