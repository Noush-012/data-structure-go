package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}
type BST struct {
	root *Node
}

// insertion in binary search tree

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

// Search key in Binary search Tree

func (b *BST) contain(key int) bool {
	currNode := b.root
	for currNode != nil {
		if key < currNode.data {
			currNode = currNode.left
		} else if key > currNode.data {
			currNode = currNode.right
		} else {
			return true
		}
	}
	return false
}
func validateBST(root *Node, min int, max int) bool {
	if root == nil {
		return true

	}

	if root.data <= min || root.data >= max {
		return false
	}

	return validateBST(root.left, min, root.data) && validateBST(root.right, root.data, max)
}

// find count of nodes in BST

func (b *BST) count(p *Node) int {
	if p != nil {
		x := b.count(p.left)
		y := b.count(p.right)
		return x + y + 1
	}
	return 0
}

// PreOrder traversal
func (node *Node) printPreOrder() {
	if node != nil {
		fmt.Printf("%d-->", node.data)
		node.left.printPreOrder()
		node.right.printPreOrder()
	}
}

// InOrder traversal
func (node *Node) printInOrder() {
	if node != nil {

		node.left.printInOrder()
		fmt.Printf("%d-->", node.data)
		node.right.printInOrder()
	}
}
func (node *Node) printPostOrder() {
	if node != nil {

		node.left.printPostOrder()
		node.right.printPostOrder()
		fmt.Printf("%d-->", node.data)
	}
}

// Remove node from BST iterative method
func (b *BST) remove(data int) {
	b.removeHelper(data, b.root, nil)
}

func (b *BST) removeHelper(data int, currentnode *Node, parentnode *Node) {
	for currentnode != nil {
		if data < currentnode.data {
			parentnode = currentnode
			currentnode = currentnode.left
		} else if data > currentnode.data {
			parentnode = currentnode
			currentnode = currentnode.right
		} else {
			if currentnode.left != nil && currentnode.right != nil {
				currentnode.data = b.getMinvalue(currentnode.right)
				b.removeHelper(currentnode.data, currentnode.right, currentnode)
			} else {
				if parentnode == nil {
					if currentnode.right == nil {
						b.root = currentnode.left
					} else {
						b.root = currentnode.right
					}
				} else {
					if parentnode.left == currentnode {
						if currentnode.right == nil {
							parentnode.left = currentnode.left
						} else {
							parentnode.left = currentnode.right
						}
					} else {
						if currentnode.right == nil {
							parentnode.right = currentnode.left
						} else {
							parentnode.right = currentnode.right
						}
					}
				}
				break
			}
		}
	}
}
func (b *BST) getMinvalue(currentnode *Node) int {
	if currentnode.left == nil {
		return currentnode.data
	}
	return b.getMinvalue(currentnode.left)
}

// Remove node from BST recursive method

// func (b *BST) Delete(key int) {
// 	deleteHelper(b.root, key)
// }

// func deleteHelper(root *Node, key int) *Node {
// 	if root == nil {
// 		return nil
// 	}
// 	if key < root.data {
// 		root.left = deleteHelper(root.left, key)
// 	} else if key > root.data {
// 		root.right = deleteHelper(root.right, key)
// 	} else {
// 		if root.left == nil {
// 			return root.right
// 		} else if root.right == nil {
// 			return root.left
// 		}
// 		root.data = findMin(root.right)
// 		root.right = deleteHelper(root.right, root.data)
// 	}
// 	return root
// }

// func findMin(node *Node) int {
// 	min := node.data
// 	for node.left != nil {
// 		min = node.left.data
// 		node = node.left
// 	}
// 	return min
// }

func main() {
	tree := &BST{}
	tree.insert(50)
	tree.insert(30)
	tree.insert(40)
	tree.insert(70)
	tree.insert(-20)
	tree.insert(10)
	tree.insert(2)
	tree.insert(100)
	tree.insert(90)
	tree.insert(22)
	// tree.Delete(30)
	// tree.Delete(50)
	tree.remove(50)
	tree.remove(40)
	fmt.Println("Pre Order")
	tree.root.printPreOrder()
	// fmt.Println(tree.contain(-20))
	// fmt.Println("In Order")
	// tree.root.printInOrder()
	// fmt.Println(tree.contain(-10))
	fmt.Println("Post Order")
	tree.root.printPostOrder()
	// count := tree.count(tree.root)
	// fmt.Println(count)

	if validateBST(tree.root, -1000, 1000) {
		fmt.Println("Valid BST")
	} else {
		fmt.Println("Invalid BST")
	}

}
