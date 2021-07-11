package main

import (
	"fmt"
	"log"
	"os"
)

const LOGFILE = "/tmp/mGo.log"

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	cLog := log.New(f, "customLogLineNumber ", log.LstdFlags)

	cLog.SetFlags(log.LstdFlags | log.Lshortfile)
	cLog.Println("Hello there!")
	cLog.Println("Another log entry!")
}
