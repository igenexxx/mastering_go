package main

import (
	"fmt"
	"os"
	"strconv"
)

func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return
}

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return min, max
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("The program needs at least 2 args!")
		return
	}

	a1, _ := strconv.Atoi(args[0])
	a2, _ := strconv.Atoi(args[1])

	fmt.Println(minMax(a1, a2))
	min, max := minMax(a1, a2)
	fmt.Println(min, max)

	fmt.Println(namedMinMax(a1, a2))
	min, max = namedMinMax(a1, a2)
	fmt.Println(min, max)
}
