package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type linkelist struct {
	head *Node
}

func (list *linkelist) add(num int) {
	newNode := &Node{
		value: num,
		next:  nil,
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

func (list *linkelist) display() {

	temp := list.head

	for temp != nil {
		fmt.Printf("%d ->", temp.value)
		temp = temp.next
	}

}

func (list *linkelist) removeDuplicates() {
	if list.head == nil {
		fmt.Println("list is emty")
		return
	}
	temp := list.head

	for temp.next != nil {
		if temp.value == temp.next.value {
			temp.next = temp.next.next
		} else {
			temp = temp.next
		}

	}
}

func main() {
	list := linkelist{}
	list.add(1)
	list.add(2)
	list.add(4)
	list.add(5)
	list.add(5)

	list.removeDuplicates()
	list.display()
}
