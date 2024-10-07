package main

import "fmt"

func main() {
	que := Queue{}
	que.display()
}

type Queue struct {
	content []int
}

func (que *Queue) display() {
	fmt.Println(que.content)
}
