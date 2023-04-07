package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}
type BST struct {
	root *Node
}

func (bst *BST) Insert(data int) {
	//fist check the root is nill then create root
	if bst.root == nil {
		bst.root = &Node{data: data}
		return
	}

	//head is not nil then traverse throgh the binary search tree and add on it
	currentNode := bst.root

	for { // in here value is equal to currentNode then add new node on its right
		if data < currentNode.data { // data less than currentNode then check left
			if currentNode.left == nil { // and if its left is nill then create a node at left
				currentNode.left = &Node{data: data}
				return
			} // left is not nil move to left
			currentNode = currentNode.left
		} else { // data is greater than currentNode the check right
			if currentNode.right == nil { // if right is nil then create a node at right
				currentNode.right = &Node{data: data}
				return
			} // rigth is not nil then move to right
			currentNode = currentNode.right
		}
	}
}

func (b *BST) delete(value int) {

	if b.root == nil {
		fmt.Println("Binary search tree is empty")
		return
	}

	currNode := b.root
	for currNode != nil {
		// To get the position to delete the node
		if value < currNode.data {
			if currNode.left != nil && currNode.data == value {
				currNode.left = removeHelper(currNode.left)
				return
			}
		}
		if value > currNode.data {
			if currNode.right != nil && currNode.data == value {
				currNode.right = removeHelper(currNode.right)
				return
			}
		} else {
			b.root = removeHelper(b.root)
			return
		}
	}
	fmt.Println("Can't find the value to delete!!!")

}

func removeHelper(removeNode *Node) *Node {
	removeNodeLeft := removeNode.left
	removeNodeRight := removeNode.right

	if removeNode.right == nil {
		return removeNodeLeft
	} else if removeNode.right.left == nil {
		removeNodeRight.left = removeNodeLeft
		return removeNodeRight
	}
	prevNode := removeNodeRight
	currNode := removeNodeRight.left
	for currNode.left != nil { // loop until find last left node
		prevNode = currNode      // find last prev node
		currNode = currNode.left // last node
	}
	//connect last left nodes right to previous node's left
	// chance last node have right node is present
	prevNode.left = currNode.right

	// now the currentNode is last left node
	// so connect last node's left as removeNode's left and right as removeNode's right
	currNode.left = removeNodeLeft
	currNode.right = removeNodeRight

	return currNode
}
func (bst *BST) Inorder() {
	if bst.root == nil {
		fmt.Println("binary search tree is empty")
		return
	}
	InorderHelper(bst.root)
	fmt.Println()
}

// func to help recursion on Inorder
func InorderHelper(node *Node) {

	if node == nil {
		return
	}

	//go to left
	InorderHelper(node.left)
	// print the node value at last
	fmt.Print(node.data, " ")
	// go to right
	InorderHelper(node.right)
}

// func to print bst values in preOrder
// root -> left -> right
func (bst *BST) PreOrder() {
	if bst.root == nil {
		fmt.Println("binary search tree is empty")
		return
	}
	preOrderHelper(bst.root)
	fmt.Println()
}

// func to help recursion for preOrder
func preOrderHelper(node *Node) {

	if node == nil {
		return
	}

	//print the node value
	fmt.Print(node.data, " ")
	//then go to left
	preOrderHelper(node.left)
	//then go to right
	preOrderHelper(node.right)
}

// print values in postOrder
// left -> right -> root

func (bst *BST) PostOrder() {
	if bst.root == nil {
		fmt.Println("binary search tree is empty")
		return
	}
	postOrderHelper(bst.root)
	fmt.Println()
}

// func to help recursion for postOrder
func postOrderHelper(node *Node) {
	if node == nil {
		return
	}

	// go to left
	postOrderHelper(node.left)
	// then go to left
	postOrderHelper(node.right)
	//last print value
	fmt.Print(node.data, " ")
}
func main() {
	b := &BST{}

	b.Insert(5)
	b.Insert(3)
	b.Insert(78)
	b.Insert(1)
	b.Insert(45)
	b.Insert(53)
	b.Insert(33)
	b.delete(78)
	// b.delete(78)

	b.Inorder()

}
