package main

import "fmt"

type linkelist struct {
	number int
	next   *linkelist
}

func (node *linkelist) lastinsert(num int) {

	var temp = &linkelist{}
	temp.number = num
	temp.next = nil
	if node == nil {
		node = temp
	} else {
		var p = &linkelist{}
		p = node

		for p.next != nil {
			p = p.next
		}
		p.next = temp
	}

}

func (node *linkelist) display() {
	var p = &linkelist{}
	p = node.next
	for p != nil {
		fmt.Printf("%d ->", p.number)
		p = p.next
	}
}

func main() {
	head := new(linkelist)
	head.lastinsert(12)
	head.lastinsert(23)
	head.lastinsert(82)
	head.lastinsert(9)

	head.display()
}
