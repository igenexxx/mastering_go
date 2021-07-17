package main

import "fmt"

func changeArr(arr [3]int) {
	arr[1] = 5
}

func changeSlice(slice []int) {
	slice[1] = 5
}

func main() {
	testArray := [3]int{1, 2, 3}

	changeArr(testArray)
	fmt.Println(testArray)

	testSlice := []int{1, 2, 3}
	changeSlice(testSlice)
	fmt.Println(testSlice)
}
