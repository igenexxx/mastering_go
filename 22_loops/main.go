package main

import "fmt"

func main() {
	k := 0
	for k < 100 {
		k++
		fmt.Println("k:", k)
	}

	// First example
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}
		if i == 95 {
			break
		}
		fmt.Print(i, " ")
	}

	// Second example
	fmt.Println()
	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}

	fmt.Println()

	// Third example
	i = 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}

		fmt.Print(i, " ")
		i++
	}

	fmt.Println()

	// Fourth example
	anArray := [5]int{0, 1, -1, 2, 3}
	for i, value := range anArray {
		fmt.Println("index:", i, "value", value)
	}
}
