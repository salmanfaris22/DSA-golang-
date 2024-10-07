package main

import "fmt"

func main() {
	s := Stack{}
	s.Push(2)
	s.Push(12)
	s.Push(10)
	s.Display()
}

type Stack struct {
	content []int
}

func (stack *Stack) Display() {
	fmt.Println(stack.content)
}
func (stack *Stack) Push(num int) {
	stack.content = append(stack.content, num)
}
