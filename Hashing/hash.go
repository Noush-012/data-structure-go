package main

import "fmt"

const size = 15

type Node struct {
	Key   int
	Value string
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func hashFunction(i, size int) int {
	return i % size
}

func insertNode(hash *HashTable, key int, value string) int {
	index := hashFunction(key, hash.Size)
	element := hash.Table[index]
	newNode := &Node{Key: key, Value: value, Next: nil}
	if element == nil {
		hash.Table[index] = newNode
	} else {
		for element.Next != nil {
			element = element.Next
		}
		element.Next = newNode
	}
	return index
}

func retrieveNode(hash *HashTable, key int) (string, bool) {
	index := hashFunction(key, hash.Size)
	element := hash.Table[index]
	if element != nil {
		for element != nil {
			if element.Key == key {
				return element.Value, true
			}
			element = element.Next
		}
	}
	return "", false
}

func main() {
	table := make(map[int]*Node, size)
	hash := &HashTable{Table: table, Size: size}
	insertNode(hash, 1, "Apple")
	insertNode(hash, 2, "Banana")
	insertNode(hash, 3, "Carrot")
	insertNode(hash, 4, "Date")
	insertNode(hash, 5, "Eggplant")

	value, ok := retrieveNode(hash, 4)
	if ok {
		fmt.Println("Value: ", value)
	} else {
		fmt.Println("Key not found")
	}
}
