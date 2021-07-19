package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		log.Fatalln("Please provide one text file to process!")
	}

	filename := arguments[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening file %s", err)
	}
	defer f.Close()

	notAMatch := 0
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
		}

		r1 := regexp.MustCompile(`.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\] .*`)
		if r1.MatchString(line) {
			match := r1.FindStringSubmatch(line)
			d1, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err != nil {
				notAMatch++
			}

			newFormat := d1.Format(time.Stamp)
			fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			continue
		}
	}
}
