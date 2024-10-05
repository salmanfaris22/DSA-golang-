package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type linkelist struct {
	head *Node
}

func (list *linkelist) convTolist(num int) {

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

func (list *linkelist) arryConv(arr []int) {
	for _, v := range arr {
		list.convTolist(v)
	}
}

func (list *linkelist) diplay() {
	temp := list.head
	for temp != nil {
		fmt.Printf("%d ->", temp.value)
		temp = temp.next
	}
}

func main() {
	list := &linkelist{}
	arr := []int{1, 2, 3, 4, 5}
	list.arryConv(arr)
	// list.convTolist(5)
	// list.convTolist(15)
	// list.convTolist(53)
	list.diplay()

}
