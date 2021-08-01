package main

import (
	"flag"
	"fmt"
)

func main() {
	minusK := flag.Bool("k", true, "k flag")
	minusO := flag.Int("o", 1, "o")
	flag.Parse()

	valueK := *minusK
	valueO := *minusO
	valueO++

	fmt.Println("-k", valueK)
	fmt.Println("-o", valueO)
}
