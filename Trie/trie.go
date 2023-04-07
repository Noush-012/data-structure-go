package main

import "fmt"

// Possible char in the trie
const AlphabetSize = 26

//Node represents each node in the trie
type Node struct {
	child [26]*Node
	isEnd bool
}

// Root of trie
type Trie struct {
	root *Node
}

func initTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// insert word in trie
func (t *Trie) insert(w string) {
	length := len(w)
	currNode := t.root
	for i := 0; i < length; i++ {
		// assigning as alphabet order a = 0; b= 1, ..... z = 26
		charIndex := w[i] - 'a'
		if currNode.child[charIndex] == nil {
			currNode.child[charIndex] = &Node{}
		}
		currNode = currNode.child[charIndex]
	}
	currNode.isEnd = true
}

// Search
func (t *Trie) search(w string) bool {
	length := len(w)
	currNode := t.root
	for i := 0; i < length; i++ {
		// assigning as alphabet order a = 0; b= 1, ..... z = 26
		charIndex := w[i] - 'a'
		if currNode.child[charIndex] == nil {
			return false
		}
		currNode = currNode.child[charIndex]
	}
	if currNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	myTrie := initTrie()
	// myTrie := &Trie{}
	myTrie.insert("noushad")
	myTrie.insert("hello")
	myTrie.insert("world")

	toAdd := []string{
		"aragorn",
		"aragon",
		"eragon",
		"oregano",
		"ergon",
		"eragon",
		"oreo",
	}

	for _, v := range toAdd {
		myTrie.insert(v)
	}
	fmt.Println(myTrie.search("noushad")) // true
	fmt.Println(myTrie.search("hello"))   // true
	fmt.Println(myTrie.search("world"))   // true
	fmt.Println(myTrie.search("Hey"))     // false
}
