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
	fmt.Printf("\n")
	for temp != nil {
		fmt.Printf("%d ->", temp.value)
		temp = temp.next
	}

}

// func (list *linkelist) reverse() {

// 	var prev *Node = nil
// 	current := list.head
// 	var next *Node = nil

// 	for current != nil {
// 		next = current.next
// 		current.next = prev
// 		prev = current
// 		current = next
// 	}
// 	list.head = prev

// }

func reversNode(list *linkelist) {
	var prev *Node = nil
	temp := list.head
	var next *Node = nil
	for temp != nil {
		next = temp.next
		temp.next = prev
		prev = temp
		temp = next
	}
	list.head = prev

}

func main() {
	list := linkelist{}
	list.add(1)
	list.add(2)
	list.add(4)
	list.add(5)
	list.add(5)
	list.display()
	reversNode(&list)
	// list.reverse()
	list.display()
}
