package main

import "fmt"

func f1(c chan int, x int) {
	fmt.Println(x)
	c <- x
}

func f2(c chan<- int, x int) {
	fmt.Println(x)
	c <- x
}

func f3(out chan<- int64, in <-chan int64) {
	fmt.Println(x)
	c <- x
}

func main() {

}
