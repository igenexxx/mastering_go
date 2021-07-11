package main

import (
	"io"
	"log"
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

	io.WriteString(os.Stdout, "This is Standard output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
	log.Fatalln("AAAA")
}
