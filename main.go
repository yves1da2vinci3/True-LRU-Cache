package main

import (
	"fmt"
)

const SIZE = 5

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Lenght int
}

type Hash map[string]*Node

type Cache struct {
	Hash  Hash
	Queue Queue
}

func NewCache() Cache {
	return Cache{Queue: newQueue(), Hash: Hash{}}
}

func newQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head
	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(str string) {
	node := &Node{}
	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Value: str}
	}
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Add(node *Node) {
	fmt.Printf("adding node %s\n", node.Value)
	temp := c.Queue.Head.Right

	c.Queue.Head.Right = node
	node.Left = c.Queue.Head

	node.Right = temp
	temp.Left = node
	c.Queue.Lenght++
	if c.Queue.Lenght > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Remove(node *Node) *Node {
	fmt.Printf("removing %s\n", node.Value)
	left := node.Left
	right := node.Right

	right.Left = left
	left.Right = right

	c.Queue.Lenght -= 1
	delete(c.Hash, node.Value)
	return node
}

func (c *Cache) Display() {
	c.Queue.Display()
}
func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d -[", q.Lenght)
	for i := 0; i < q.Lenght; i++ {
		fmt.Printf("{%s}", node.Value)
		if i < q.Lenght-1 {
			fmt.Printf("<->")
		}
		node = node.Right

	}
	fmt.Println("]")
}

func main() {
	fmt.Println("Starting Caching")
	cache := NewCache()
	for _, word := range []string{"parrot", "avocado", "dragonfruit", "tree", "potato", "tree", "dog"} {
		cache.Check(word)
		cache.Display()
	}
}
