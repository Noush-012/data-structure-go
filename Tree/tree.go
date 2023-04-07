package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}
type BST struct {
	root *Node
}

func (b *BST) insert(data int) {
	newNode := new(Node)
	newNode.data = data
	if b.root == nil {
		b.root = newNode
	} else {
		currNode := b.root
		for currNode != nil {
			if newNode.data < currNode.data {
				if currNode.left == nil {
					currNode.left = newNode
					return
				} else {
					currNode = currNode.left
				}
			} else {
				if currNode.right == nil {
					currNode.right = newNode
					return
				} else {
					currNode = currNode.right
				}
			}

		}
	}
}

func (node *Node) preOrder() {
	if node != nil {
		fmt.Printf("%d-->", node.data)
		node.left.preOrder()
		node.right.preOrder()
	}
}

func main() {
	b := BST{}

	b.insert(50)
	b.insert(30)
	b.insert(20)
	b.insert(40)
	b.insert(80)
	b.insert(10)
	fmt.Println("Pre Order")
	b.root.preOrder()

}
