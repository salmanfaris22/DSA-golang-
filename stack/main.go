package main

import "fmt"

func main() {
	s := Stack{}
	s.Push(2)
	s.Push(12)
	s.Push(10)
	deleted := s.Pop()
	fmt.Println("items delited :", deleted)
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
func (stack *Stack) Pop() int {
	lastIndex := len(stack.content) - 1
	lastItem := stack.content[lastIndex]
	stack.content = stack.content[:lastIndex]

	return lastItem
}
