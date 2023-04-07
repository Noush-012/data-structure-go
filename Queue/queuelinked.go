package main
type Node struct{
	data int
	next *Node
}
type queue struct{
	first *Node
	last *Node
}

func (q *queue) equeue(data int){
	newNode :=&Node{data, nil}
	if q.last == nil{
		q.first = newNode
		q.last = newNode
	}else{
		q.last.next = newNode
		q.last  = newNode
	}
}
func main() {
	
}