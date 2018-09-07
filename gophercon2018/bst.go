package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}
	if n.key < key {
		return n.right.Search(key)
	} else if n.key > key {
		return n.left.Search(key)
	}
	return true
}

func (n *Node) Insert(key int) {
	if n.key < key {
		if n.right == nil {
			n.right = &Node{key: key}
		} else {
			n.right.Insert(key)
		}
	} else if n.key > key {
		if n.left == nil {
			n.left = &Node{key: key}
		} else {
			n.left.Insert(key)
		}
	}
}

func (n *Node) Min() int {
	// TODO:
	if n.left == nil {
		return n.key
	} else {
		return n.left.Min()
	}
}

func (n *Node) Max() int {
	if n.right == nil {
		return n.key
	} else {
		return n.right.Min()
	}
}

func (n *Node) Delete(key int) *Node {
	if n.key < key {
		n.right = n.right.Delete(key)
	} else if n.key > key {
		n.left = n.left.Delete(key)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		min := n.right.Min()
		n.key = min
		n.right = n.right.Delete(min)
	}
	return n
}

func main() {
	tree := &Node{key: 6}
	tree.Insert(5)
	fmt.Printf("head node %d\n", tree.key)
	fmt.Printf("min %d\n", tree.Min())
	fmt.Printf("max %d\n", tree.Max())
}
