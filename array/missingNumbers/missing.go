package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 4}
	miss := missing(arr)
	fmt.Println(miss)
}

func missing(arr []int) int {
	n := len(arr) + 1
	sum := n * (n + 1) / 2
	num := 0
	for i := 0; i < len(arr); i++ {
		num += arr[i]
	}
	return sum - num
}
