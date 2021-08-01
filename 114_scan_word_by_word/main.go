package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func wordByWord(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: byWord <file1> [<file2> ...]\n")
		return
	}

	for _, filePath := range flag.Args() {
		wordByWord(filePath)
	}
}
