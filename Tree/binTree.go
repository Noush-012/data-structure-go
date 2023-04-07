package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert adds a new node to the binary tree
func (n *Node) Insert(value int) *Node {
	if n == nil {
		return &Node{Value: value}
	}

	if value < n.Value {
		n.Left = n.Left.Insert(value)
	} else {
		n.Right = n.Right.Insert(value)
	}

	return n
}

// Search searches for a node with a specific value in the binary tree
func (n *Node) Search(value int) *Node {
	if n == nil || n.Value == value {
		return n
	}

	if value < n.Value {
		return n.Left.Search(value)
	} else {
		return n.Right.Search(value)
	}
}
func (n *Node) InorderTraversal() {
	if n == nil {
		return
	}

	n.Left.InorderTraversal()
	fmt.Println(n.Value)
	n.Right.InorderTraversal()
}

// To delete a node from a binary tree in Golang, you need to perform three steps:

// 1. Find the node to be deleted and its parent node.
// 2. Determine the replacement node for the deleted node.
// 3. Update the parent node to point to the replacement node.

func (n *Node) Delete(value int) *Node {
	// Find the node to be deleted and its parent node
	var parent *Node
	node := n
	for node != nil && node.Value != value {
		parent = node
		if value < node.Value {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	// If the node was not found, return the original tree
	if node == nil {
		return n
	}

	// Determine the replacement node
	var replacement *Node
	if node.Left == nil && node.Right == nil {
		replacement = nil
	} else if node.Left == nil {
		replacement = node.Right
	} else if node.Right == nil {
		replacement = node.Left
	} else {
		replacement = node.Right
		for replacement.Left != nil {
			replacement = replacement.Left
		}
		replacement.Right = node.Right
		replacement.Left = node.Left
	}

	// Update the parent node to point to the replacement node
	if parent == nil {
		return replacement
	} else if parent.Left == node {
		parent.Left = replacement
	} else {
		parent.Right = replacement
	}

	return n
}

func (n *Node) LevelOrderTraversal() {
	if n == nil {
		return
	}

	queue := []*Node{n}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
func main() {
	// Create a new binary tree with the root node having a value of 5
	root := &Node{Value: 5}

	// Insert some nodes into the binary tree
	root.Insert(3)
	root.Insert(7)
	root.Insert(1)
	root.Insert(9)
	root.Insert(4)

	// Search for a node with a value of 4 in the binary tree
	node := root.Search(4)
	if node != nil {
		fmt.Println("Found node with value:", node.Value)
	} else {
		fmt.Println("Node not found")
	}

	root.InorderTraversal()
	fmt.Println("Visual level order")
	root.LevelOrderTraversal()
}
