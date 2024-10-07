package main

import (
	"errors"
	"fmt"
)

func main() {
	s := Stack{}
	emty := s.IsEmpty()
	if emty {
		fmt.Println("stack is mety")
	} else {
		fmt.Println("stack is not emty")
	}

	s.Push(2)
	s.Push(12)
	s.Push(10)
	size := s.Size()
	fmt.Println("stack suze is :", size)
	top, err := s.Peek()
	if err != nil {
		fmt.Println(err.Error())

	} else {
		fmt.Println("top element :", top)
	}
	deleted, err := s.Pop()
	if err != nil {
		fmt.Println(err.Error())

	} else {
		fmt.Println("items delited :", deleted)
	}
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
func (stack *Stack) Pop() (int, error) {
	if len(stack.content) == 0 {
		return 0, errors.New("stack is emty")
	}
	lastIndex := len(stack.content) - 1
	lastItem := stack.content[lastIndex]
	stack.content = stack.content[:lastIndex]

	return lastItem, nil
}
func (stack *Stack) Peek() (int, error) {
	if len(stack.content) == 0 {
		return 0, errors.New("stack is emty")
	}
	return stack.content[len(stack.content)-1], nil
}
func (stack *Stack) IsEmpty() bool {
	if len(stack.content) != 0 {
		return false
	} else {
		return true
	}

}
func (stack *Stack) Size() int {
	return len(stack.content)
}
