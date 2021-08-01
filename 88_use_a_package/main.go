package main

import (
	"fmt"
	a_package "github.com/igenex/mastering_go/87_a_package"
)

func main() {
	fmt.Println("Using a_package!")
	a_package.A()
	a_package.B()
	fmt.Println(a_package.MyConstant)
}
