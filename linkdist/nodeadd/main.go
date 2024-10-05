package main

import (
	"fmt"
)

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

func (list *linkelist) deliteLast() {

	if list.head == nil {
		fmt.Println("pleas enter node")

	} else if list.head.next == nil {
		list.head = nil

	} else {
		temp := list.head

		for temp.next.next != nil {

			temp = temp.next
		}
		temp.next = nil

	}

}

func (list *linkelist) delistspeceficNode(num int) {

	if list.head == nil {
		fmt.Println("pleas enter node")
		return
	}

	if list.head.number == num {
		list.head = list.head.next
		return
	}

	temp := list.head
	for temp.next.number != num {
		temp = temp.next

		if temp.next == nil {
			fmt.Println("\n", num, "this number not find in this linkedlist", "\n")
			return
		}
	}
	fmt.Println(temp)
	temp.next = temp.next.next

}

func (list *linkelist) deletfirst() {

	if list.head == nil {
		fmt.Println("pleas enter node")

	} else {
		temp := list.head
		list.head = temp.next
	}

}

func (list *linkelist) display() {
	fmt.Printf("\n")
	temp := list.head
	for temp != nil {
		fmt.Printf("%d ->", temp.number)
		temp = temp.next
	}

}

func (list *linkelist) reversedisplay(node *Node) {

	if node == nil {
		return
	}
	list.reversedisplay(node.next)
	fmt.Printf("%d ->", node.number)

}

func main() {
	list := linkelist{}
	list.insertLast(3)
	list.insertLast(30)
	list.insertLast(13)
	list.insertLast(10)
	list.insertLast(2)
	// list.insertFirst(1011)
	list.display()
	// list.delistspeceficNode(120)
	fmt.Printf("\n")
	list.reversedisplay(list.head)
	// list.deliteLast()
	// list.display()
	// list.deletfirst()
	// list.display()

}
