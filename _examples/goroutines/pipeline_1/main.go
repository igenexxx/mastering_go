package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

var numberOfPhrases int

func readLineByLine(file string, c chan<- string, wg *sync.WaitGroup) {
	f, err := os.Open(file)
	if err != nil {
		close(c)
	}
	defer f.Close()

	r := bufio.NewScanner(f)
	for r.Scan() {
		wg.Add(1)
		c <- r.Text()
	}

	if err := r.Err(); err != nil {
		close(c)
	}

	close(c)
}

func phraseCounter(phrase string, in <-chan string, wg *sync.WaitGroup) {
	for line := range in {
		if strings.Contains(strings.ToLower(line), strings.ToLower(phrase)) {
			numberOfPhrases++
		}
		wg.Done()
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: byLine <file1> [<file2> ...]\n")
		return
	}
	var waitGroup sync.WaitGroup

	for _, file := range flag.Args() {
		fileChan := make(chan string)
		go readLineByLine(file, fileChan, &waitGroup)
		phraseCounter("qua", fileChan, &waitGroup)
	}

	fmt.Println("Phrase is faced", numberOfPhrases, "times")
}
