package main

import "fmt"

func main() {
	fmt.Println(InsertingSort([]int{1, 6, 7, 10, 5, 2, 3, 4}))
}

func InsertingSort(arr []int) []int {
	for i := 0; i < len(arr); i++ { // 1, 6, 7, 10, 5, 2, 3, 4

		key := arr[i]                // key=2
		j := i - 1                   // j= 5
		for j >= 0 && arr[j] > key { //j>=0 && arr[5](10) > 2

			arr[j+1] = arr[j] // // arr[]= 6
			j--               //j-- j=0
		}

		arr[j+1] = key //arr[j+1][1+1]=key
		// 1, 5,6,7,7,10
	}
	return arr
}
