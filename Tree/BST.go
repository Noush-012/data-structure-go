package main

import "fmt"

var count int = 1
var sCount int

type Node struct {
	data        int
	left, right *Node
}

func (n *Node) insert(data int) {
	if n.data < data {
		// Move right
		if n.right == nil {
			n.right = &Node{data: data}
			count++
		} else {
			n.right.insert(data)
		}
	}
	if n.data > data {
		// Move left
		if n.left == nil {
			n.left = &Node{data: data}
			count++
		} else {
			n.left.insert(data)
		}
	}
}

func (n *Node) search(key int) (bool, int) {
	sCount++
	if n == nil {
		return false, sCount
	}
	if n.data < key {
		// Search right
		return n.right.search(key)

	} else {
		if n.data > key {
			// Search left

			return n.left.search(key)

		}
	}
	return true, sCount
}

// func (n *Node) delete(key int) {
// 	if n == nil {
// 		fmt.Println("Tree empty")
// 	} else if key < n.data {
// 		n.left.delete(key)
// 	}
// }

func (n *Node) delete(key int) *Node {
	if n == nil {
		fmt.Println("Tree empty")
		return n
	}

	if key < n.data {
		n.left = n.left.delete(key)
	} else if key > n.data {
		n.right = n.right.delete(key)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		// Node to be deleted has two children
		minRight := n.right.min()
		n.data = minRight.data
		n.right = n.right.delete(minRight.data)
	}

	return n
}
func (n *Node) min() *Node {
	if n.left == nil {
		return n
	}
	return n.left.min()
}

func main() {

	root := &Node{data: 50}
	root.insert(100)
	root.insert(30)
	root.insert(20)
	root.insert(40)
	root.insert(80)
	root.insert(60)
	root.delete(20)

	// fmt.Println(root.search(50))
	// fmt.Println("Total tree:", count)

}
