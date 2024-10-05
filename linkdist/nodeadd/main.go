package main

import "fmt"

type Node struct {
	number int
	next   *Node
}

type linkelist struct {
	head *Node
}

func (list *linkelist) insertLast(num int) {
	newNode := &Node{
		number: num,
		next:   nil,
	}
	if list.head == nil {
		list.head = newNode
	} else {
		temp := list.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
	}

}

func (list *linkelist) insertFirst(num int) {

	newNode := &Node{
		number: num,
		next:   list.head,
	}
	list.head = newNode
}

func (list *linkelist) display() {
	temp := list.head
	for temp != nil {
		fmt.Printf("%d ->", temp.number)
		temp = temp.next
	}
}

func main() {
	list := linkelist{}
	list.insertLast(3)
	list.insertLast(10)
	list.insertLast(15)
	list.insertFirst(300)
	list.insertFirst(10)
	list.display()
}
