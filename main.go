package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type linkelist struct {
	head *Node
}

func (list *linkelist) insert(num int) {

	newNode := &Node{
		data: num,
		next: nil,
	}
	if list.head == nil {
		list.head = newNode
	} else {
		newNode.next = list.head
		list.head = newNode
	}

}

func (list *linkelist) display() {

	temp := list.head

	for temp != nil {
		fmt.Printf("%d->", temp.data)
		temp = temp.next
	}

}

func main() {
	list := linkelist{}
	list.insert(4)
	list.insert(23)
	list.insert(411)
	list.insert(10)
	list.display()

}
