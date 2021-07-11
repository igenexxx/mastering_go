package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		log.Fatalln("Please give one or more floats.")
	}

	min, _ := strconv.ParseFloat(arguments[0], 64)
	max := min

	for i := 1; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
