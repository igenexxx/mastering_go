package main

import "fmt"

func returnPointer(x int) *int {
	y := x * x
	return &y
}

func main() {
	sq := returnPointer(10)
	fmt.Println("sq value:", *sq)
	fmt.Println("sq memory address:", sq)
}
