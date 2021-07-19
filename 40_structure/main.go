package main

import "fmt"

type aStructrue struct {
	person string
	height int
	weight int
}

func main() {
	p1 := aStructrue{"fmt", 12, -2}
	p2 := aStructrue{weight: 12, height: -2}

	fmt.Println(p1, p2)
}
