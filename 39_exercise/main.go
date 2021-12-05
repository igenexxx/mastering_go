package main

import (
	"fmt"
)

func main() {
	const (
		mo = iota + 1
		tu
		we
		th
		fr
		sa
		su
	)

	fmt.Println(we)

	test := []string{"hi", "hello", "alloha", "привет", "здраво"}
	testMap := make(map[int]string)

	for i, s := range test {
		testMap[i] = s
	}

	fmt.Println(testMap)

	const (
		a_4 = iota << 4
		b_4
		c_4
		d_4
	)

	fmt.Println(a_4, b_4, c_4, d_4)
}
