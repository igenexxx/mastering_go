package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var numberOfPhrases int

func readLineByLine(file string, c chan<- string) {
	f, err := os.Open(file)
	if err != nil {
		close(c)
	}
	defer f.Close()

	r := bufio.NewScanner(f)
	for r.Scan() {
		fmt.Println(r.Text())
		c <- r.Text()
	}

	if err := r.Err(); err != nil {
		close(c)
	}

	close(c)
}

func phraseCounter(phrase string, in <-chan string) {
	for line := range in {
		numberOfPhrases += strings.Count(strings.ToLower(line), strings.ToLower(phrase))
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: byLine <file1> [<file2> ...]\n")
		return
	}

	for _, file := range flag.Args() {
		fileChan := make(chan string)
		go readLineByLine(file, fileChan)
		phraseCounter("qua", fileChan)
	}

	fmt.Println("Phrase is faced", numberOfPhrases, "times")
}
