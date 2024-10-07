package main

import "fmt"

func main() {

	arr := []int{1, 2, 3}
	arr1 := []int{4, 5, 6}

	newArray := mergaArray(arr, arr1)

	fmt.Println(newArray)
}

func mergaArray(arr, arr1 []int) []int {

	newarr := arr
	for i := 0; i < len(arr1); i++ {
		
	}
	return newarr
}
