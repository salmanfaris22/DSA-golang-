package main

import (
	"errors"
	"fmt"
)

func main() {
	que := Queue{}
	que.Enqueue(2)
	que.Enqueue(210)
	// item, err := que.Dequeue()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println("poped success item is:", item)
	// }

	que.Enqueue(12)
	top, _ := que.Peek()
	fmt.Println("top element", top)
	que.display()
}

type Queue struct {
	content []int
}

func (que *Queue) display() {
	fmt.Println(que.content)
}

func (que *Queue) Enqueue(num int) {
	que.content = append(que.content, num)
}

func (que *Queue) Dequeue() (int, error) {
	if len(que.content) == 0 {
		return 0, errors.New("queue is emty")
	}
	item := que.content[0]
	que.content = que.content[1:]
	return item, nil
}

func (que *Queue) Peek() (int, error) {
	if len(que.content) == 0 {
		return 0, errors.New("queue is emty")
	}

	return que.content[0], nil
}
func (que *Queue) IsEmty() bool {

	return len(que.content) == 0
}
func (que *Queue) Size() int {
	return len(que.content)
}
