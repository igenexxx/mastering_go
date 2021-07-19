package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

func findIp(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := fmt.Sprintf("%s\\.%s\\.%s\\.%s", partIP, partIP, partIP, partIP)
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		log.Fatalf("usage: %s logFile\n", filepath.Base(arguments[0]))
	}

	for _, filename := range arguments[1:] {
		f, err := os.Open(filename)
		if err != nil {
			log.Fatalf("error opening file %s\n", err)
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')

			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
				break
			}

			ip := findIp(line)
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}
			fmt.Println(ip)
		}
	}
}
