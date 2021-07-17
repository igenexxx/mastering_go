package main

import "fmt"

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
}
