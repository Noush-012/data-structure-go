package main

import "fmt"

const max = 5

type hashtable struct {
	ar [max]*linkedList
}
type linkedList struct {
	head *Node
}
type Node struct {
	data string
	next *Node
}

func Init() *hashtable {
	r := &hashtable{}
	for i := range r.ar {
		r.ar[i] = &linkedList{}
	}
	return r
}

// Hash table operations
// Insertion to the hash table array
func (hh *hashtable) Insert(key string) {
	index := hash(key)
	hh.ar[index].insert(key)
}
func (hh *hashtable) Search(key string) bool {
	index := hash(key)
	return hh.ar[index].search(key)
}
func (hh *hashtable) Delete(key string) {
	index := hash(key)
	hh.ar[index].delete(key)
}

// Linked list Operation
// create a node with key and insert the node in the linked list after index generates
func (l *linkedList) insert(key string) {
	newNode := &Node{data: key}
	newNode.next = l.head
	l.head = newNode
}

func (l *linkedList) search(k string) bool {
	// index := hash(key)
	current := l.head
	for current != nil {
		if current.data == k {
			return true
		}
		current = current.next
	}
	return false
}
func (l *linkedList) delete(k string) {
	if l.head.data == k {
		l.head = l.head.next
		fmt.Println("Deleted successfuly")
		return
	}
	prevNode := l.head
	for prevNode.next != nil {
		if prevNode.next.data == k {
			//delete
			prevNode.next = prevNode.next.next
			fmt.Println("Deleted successfuly!!!")
		}
		prevNode = prevNode.next
	}
}
func hash(key string) int {
	sum := 0
	for i := 0; i < len(key); i++ {

		sum += int(key[i])
	}
	return sum % max
}
func main() {
	h := Init()
	list := []string{
		"Apple",
		"Orange",
		"Grapes",
		"Peeches",
		"Melon",
	}
	for _, v := range list {
		h.Insert(v)
	}
	// h := Init()
	// fmt.Println(hash("AAAAAAA"))
	// lltest := &linkedList{}
	// lltest.insert("Noushad")
	// fmt.Println(h.Search("Melon"))
	h.Delete("Grapes")

	fmt.Println(h.Search("Orange"))

}
