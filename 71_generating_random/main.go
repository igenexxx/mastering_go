package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	MIN := 0
	MAX := 100
	TOTAL := 100
	SEED := time.Now().Unix()

	args := os.Args[1:]

	if len(args) <= 4 {
		fmt.Println("Usage: ./randomNumbers MIN MAX TOTAL SEED")
	}

	switch len(args) {
	case 1:
		MIN, _ = strconv.Atoi(args[0])
		MAX = MIN + 100
	case 2:
		MIN, _ = strconv.Atoi(args[0])
		MAX, _ = strconv.Atoi(args[1])
	case 3:
		MIN, _ = strconv.Atoi(args[0])
		MAX, _ = strconv.Atoi(args[1])
		TOTAL, _ = strconv.Atoi(args[2])
	case 4:
		MIN, _ = strconv.Atoi(args[0])
		MAX, _ = strconv.Atoi(args[1])
		TOTAL, _ = strconv.Atoi(args[2])
		SEED, _ = strconv.ParseInt(args[3], 10, 64)
	default:
		fmt.Println("Using default values!")
	}

	rand.Seed(SEED)
	for i := 0; i < TOTAL; i++ {
		myRand := random(MIN, MAX)
		fmt.Print(myRand)
		fmt.Print(" ")
	}
	fmt.Println()
}
