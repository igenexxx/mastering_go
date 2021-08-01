package main

import "fmt"

type first struct{}

func (f first) F() {
	f.shared()
}

func (f first) shared() {
	fmt.Println("This is shared() from first!")
}

type second struct {
	first
}

func (s second) shared() {
	fmt.Println("This is shared() from second!")
}

func main() {
	first{}.F()
	second{}.shared()
	i := second{}
	j := i.first
	j.F()
}
