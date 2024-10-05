package main

import "fmt"

type Node struct {
	prev  *Node
	value int
	next  *Node
}

type linkedlist struct {
	head *Node
	tail *Node
}

func (list *linkedlist) add(num int) {
	newNode := &Node{
		prev:  nil,
		value: num,
		next:  nil,
	}
	if list.head == nil && list.tail == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
}

func (list *linkedlist) display() {
	temp := list.head
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}
	fmt.Println("\nForword Order")
	for temp != nil {
		fmt.Printf("%d ->", temp.value)
		temp = temp.next
	}
}
func (list *linkedlist) reversedisplay() {
	temp := list.tail
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}
	fmt.Println("\nrevers Order")
	for temp != nil {
		fmt.Printf("%d ->", temp.value)
		temp = temp.prev
	}
}

func (list *linkedlist) addBeggining(num int) {
	newNode := &Node{
		prev:  nil,
		value: num,
		next:  nil,
	}
	if list.head == nil && list.tail == nil {

		list.head = newNode
		list.tail = newNode

	} else {
		list.head.prev = newNode
		newNode.next = list.head
		list.head = newNode
	}
}

func main() {
	list := &linkedlist{}
	list.add(3)
	list.add(4)
	list.add(5)
	list.add(6)
	list.addBeggining(2)
	list.addBeggining(1)
	list.display()
	list.reversedisplay()
}
