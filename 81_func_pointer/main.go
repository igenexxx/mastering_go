package main

import "fmt"

func getPointer(v *float64) float64 {
	return *v * *v
}

func main() {
	x := 12.2
	fmt.Println(getPointer(&x))
	x = 12
	fmt.Println(getPointer(&x))
}
