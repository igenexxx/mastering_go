package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

const dateRegex = `.*(\w{3}  \d\d? \d{2}:\d{2}:\d{2}).*`

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	r := regexp.MustCompile(dateRegex)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if r.MatchString(line) {
			match := r.FindStringSubmatch(line)
			date, err := time.Parse("Jan  2 15:04:05", match[1])

			if err != nil {
				fmt.Println("Not a valid date time format!")
			}

			newFormat := date.Format(time.RFC822)
			fmt.Println(newFormat)
		} else {
			fmt.Println("Not a match!")
		}
	}
}
