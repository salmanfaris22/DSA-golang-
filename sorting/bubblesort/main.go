package main

import "fmt"

func main() {

	numbers := []int{2, 1, 9, 4, 7, 6, 3}

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			Swap(numbers, i, i+1)
			i = 0
		}
	}
	fmt.Println(numbers)
}

func Swap(arr []int, num1, num2 int) {
	arr[num1], arr[num2] = arr[num2], arr[num1]
}
