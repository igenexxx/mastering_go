package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[0]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
