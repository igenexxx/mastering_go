package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: permissions filename\n")
		return
	}

	filename := args[1]
	info, _ := os.Stat(filename)
	mode := info.Mode()
	fmt.Println(filename, "mode is", mode.String()[1:10])
}
