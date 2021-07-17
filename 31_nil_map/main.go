package main

import "fmt"

func main() {
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println(aMap)

	aMap = nil
	fmt.Println(aMap)
	fmt.Println("Hi", aMap["hi"])
	aMap["test"] = 1
}
