package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type bstTree struct {
	root *Node
}

func (bst *bstTree) insert(num int) {
	newNode := new(Node)
	newNode.data = num
	if bst.root == nil {
		bst.root = newNode
	} else {
		currNode := bst.root
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

func (t *Tree) remove(data int) {
	t.root = removeHelper(data, t.root, nil)
}

func removeHelper(data int, currentNode *Node, parentNode *Node) *Node {
	if currentNode == nil {
		return nil
	}

	if data < currentNode.data {
		currentNode.left = removeHelper(data, currentNode.left, currentNode)
	} else if data > currentNode.data {
		currentNode.right = removeHelper(data, currentNode.right, currentNode)
	} else {
		if currentNode.left != nil && currentNode.right != nil {
			minNode := minValue(currentNode.right)
			currentNode.data = minNode.data
			currentNode.right = removeHelper(minNode.data, currentNode.right, currentNode)
		} else if parentNode == nil {
			if currentNode.left != nil {
				currentNode = currentNode.left
			} else {
				currentNode = currentNode.right
			}
		} else if parentNode.left == currentNode {
			if currentNode.left != nil {
				parentNode.left = currentNode.left
			} else {
				parentNode.left = currentNode.right
			}
		} else if parentNode.right == currentNode {
			if currentNode.left != nil {
				parentNode.right = currentNode.left
			} else {
				parentNode.right = currentNode.right
			}
		}
	}

	return currentNode
}

func minValue(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}

	return node
}

func main() {
	tree := &Tree{}

	tree.insert(10)
	tree.insert(8)
	tree.insert(11)
	tree.insert(50)
	tree.insert(100)
	tree.remove(50)

	fmt.Println(tree.contain(50))
}
