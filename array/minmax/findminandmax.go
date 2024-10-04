package main

import "fmt"

func main() {

	arr := []int{1, 2, 31, 4, 5}

	min, max := minandmax(arr)

	fmt.Println(min, max)
}

func minandmax(arr []int) (int, int) {
	min, max := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		} else if arr[i] > max {
			max = arr[i]
		}
	}

	return max, min
}
