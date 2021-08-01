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
	Min := 0
	Max := 94
	var Length int64 = 8

	args := os.Args[1:]

	switch len(args) {
	case 1:
		Length, _ = strconv.ParseInt(args[0], 10, 64)
	default:
		fmt.Println("Using default values!")
	}

	Seed := time.Now().Unix()
	rand.Seed(Seed)

	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(Min, Max)
		newChar := string(startChar[0] + byte(myRand))
		fmt.Print(newChar)
		if i == Length {
			break
		}
		i++
	}
	fmt.Println()
}
