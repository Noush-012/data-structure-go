package main

import "fmt"

type Node struct {
	Key   string
	Value int
	Next  *Node
}

type linkedList struct {
	head *Node
	tail *Node
}

func (list *linkedList) addNode(key string, value int) {
	newNode := new(Node)
	newNode.Value = value
	newNode.Key = key
	newNode.Next = nil

	if list.head == nil {
		list.head = newNode

	} else {
		list.tail.Next = newNode
	}

	list.tail = newNode

}
func (list *linkedList) display(index int) {
	temp := new(Node)
	if list.head == nil {
		fmt.Println("Empty list!")
	}

	temp = list.head

	for temp != nil {
		fmt.Println("[", index, "]", temp.Key, ":", temp.Value)
		temp = temp.Next
	}

}

const tableSize = 10

type HashTable struct {
	arr [tableSize]*linkedList
}

func (ht *HashTable) hash(key string) int {
	sum := 0
	for i := 0; i < len(key); i++ {

		sum += int(key[i])
	}
	return sum % tableSize
}

func (ht *HashTable) Insert(key string, value int) {

	index := ht.hash(key)

	if ht.arr[index] == nil {
		list := new(linkedList)
		ht.arr[index] = list
		list.addNode(key, value)
	} else {
		ht.arr[index].addNode(key, value)
	}
}

func (ht *HashTable) Search(key string) int {

	index := ht.hash(key)

	if ht.arr[index] == nil {

		return -1
	} else {
		temp := ht.arr[index].head
		for temp != nil {
			if temp.Key == key {
				return temp.Value
			}
			temp = temp.Next
		}
	}

	return -1
}

func (ht *HashTable) displayHashTable() {
	var i = 0
	for i < tableSize {

		if ht.arr[i] != nil {
			ht.arr[i].display(i)
		}
		i++
	}
}

func main() {

	ht := HashTable{}

	ht.Insert("Arun", 12)
	ht.Insert("rahul", 13)
	ht.Insert("gokul", 14)
	ht.Insert("ramu", 15)
	ht.Insert("Shinu", 16)
	ht.Insert("Salman", 17)

	ht.displayHashTable()

	searchWord := "akdfksd"
	search := ht.Search(searchWord)
	if search == -1 {
		fmt.Println(searchWord, " : ", "Searched key is not present!")
	} else {
		fmt.Println(searchWord, " : ", search)
	}

}
