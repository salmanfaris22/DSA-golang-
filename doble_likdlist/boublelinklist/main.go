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
		fmt.Println("\n List is empty")
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

func (list *linkedlist) delitlast() {
	if list.tail == nil {
		fmt.Println("\n pleas enter Node")
		return
	}
	if list.tail.prev == nil {
		list.head = nil
		list.tail = nil
		return
	}
	list.tail.prev.next = nil
	list.tail = list.tail.prev
}
func (list *linkedlist) deletfirst() {
	if list.head == nil {
		fmt.Println("\n pleas enter Node")
		return
	}
	if list.head.next == nil {
		list.head = nil
		list.tail = nil
		return
	}
	list.head = list.head.next
	list.head.prev = nil
}

func (list *linkedlist) feseficdeilit(num int) {
	temp := list.head
	for temp.value != num {
		temp = temp.next
		if temp == nil {
			fmt.Println("\n number cant find")
			return
		}
	}
	if temp.prev != nil {
		temp.prev.next = temp.next
	} else {
		list.head = temp.next
	}
	if temp.next != nil {
		temp.next.prev = temp.prev
	} else {
		list.tail = temp.prev
	}

}
func main() {
	list := &linkedlist{}
	list.add(3)
	list.add(4)
	list.add(5)
	list.add(6)
	list.display()
	list.feseficdeilit(3)

	// list.deletfirst()
	// list.deletfirst()
	list.reversedisplay()
	list.display()
	// list.reversedisplay()
}
