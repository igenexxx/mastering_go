package main

import "fmt"

func main() {
	// initialize
	aSliceLiteral := []int{1, 2, 3, 4, 5}
	integer := make([]int, 20)
	fmt.Println(aSliceLiteral[3])
	fmt.Println(integer[5])

	// iterate
	for i := 0; i < len(integer); i++ {
		fmt.Println(integer[i])
	}

	// empty existing slice
	aSliceLiteral = nil

	// add element
	integer = append(integer, 12345)

	firstElement := integer[0]
	lastElement := integer[len(integer)-1]
	fmt.Println(firstElement, lastElement)
	fmt.Println(integer[1:3])
}
