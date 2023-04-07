package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	}

	if value < root.Value {
		root.Left = Insert(root.Left, value)
	} else {
		root.Right = Insert(root.Right, value)
	}

	return root
}

func TraverseInorder(root *Node) {
	if root == nil {
		return
	}

	TraverseInorder(root.Left)
	fmt.Printf("%d ", root.Value)
	TraverseInorder(root.Right)
}

func main() {
	var root *Node

	root = Insert(root, 10)
	root = Insert(root, 20)
	root = Insert(root, 5)
	root = Insert(root, 15)
	root = Insert(root, 30)

	fmt.Println("Inorder traversal:")
	TraverseInorder(root)
}
