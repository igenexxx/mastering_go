package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func initRange(min, max int, c chan<- int) {
	for i := min; i <= max; i++ {
		c <- i
	}
	close(c)
}

func makeSquareOfNaturals(out chan<- int, in <-chan int) {
	for num := range in {
		out <- int(math.Pow(2, float64(num)))
	}
	close(out)
}

func sumOfSquares(in <-chan int) {
	var sum int
	for x := range in {
		sum += x
	}
	fmt.Printf("The sum of the range is %d\n", sum)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters!")
		return
	}

	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d\n", n1, n2)
		return
	}

	A := make(chan int)
	B := make(chan int)

	go initRange(n1, n2, A)
	go makeSquareOfNaturals(B, A)
	sumOfSquares(B)
}
