package main

import "fmt"

type Digit int
type Power2 int

const Pi = 3.1415926

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func even(input int) int {
	if input%2 == 0 {
		return input
	}

	return input + 1
}

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(Pi)

	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)
	fmt.Println(One)
	fmt.Println(Two)

	const (
		p2_0 Power2 = 1 << iota
		p2_2
		p2_4
		p2_6
	)

	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
	fmt.Println("2^4:", p2_4)
	fmt.Println("2^6:", p2_6)

	const (
		a1 int = iota % (4)
		a2
		a3
		a4
		a5
		a6
		a7
		a8
		a9
	)

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println("a4:", a4)
	fmt.Println("a5:", a5)
	fmt.Println("a6:", a6)
	fmt.Println("a7:", a7)
	fmt.Println("a8:", a8)
	fmt.Println("a9:", a9)
}
