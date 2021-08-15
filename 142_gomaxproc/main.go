package main

import (
	"fmt"
	"runtime"
)

func getGoMaxProcs() int {
	return runtime.GOMAXPROCS(0)
}

func main() {
	fmt.Printf("GOMAXPROCS: %d\n", getGoMaxProcs())
}
