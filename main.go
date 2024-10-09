package main

import "fmt"

type Node struct {
	next *Node

	val int
}

type linkedlist struct {
	head *Node
}

func (list *linkedlist) Push(num int) {
	temp := list.head

	newNode := &Node{
		val: num,
	}
	if temp == nil {
		temp = newNode

	}
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode

}

func display(head *Node) {
	temp := head
	for temp != nil {
		fmt.Println("lkmk")
		fmt.Printf("%d ->", temp.val)
		temp = temp.next
	}
}

var ll = linkedlist{}

func arrayToLinkedList(arr []int) *Node {

	for _, v := range arr {
		ll.Push(v)
	}

	return ll.head
}

func main() {

	arr := []int{1, 2, 3, 4, 5}
	head := arrayToLinkedList(arr)

	display(head)
}
