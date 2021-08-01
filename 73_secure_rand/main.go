package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func generatePass(s int64) (string, error) {
	b, err := generateBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func main() {
	var Length int64 = 8
	args := os.Args[1:]
	switch len(args) {
	case 1:
		Length, _ = strconv.ParseInt(args[0], 10, 64)
		if Length <= 0 {
			Length = 8
		}
	default:
		fmt.Println("Using default values!")
	}

	myPass, err := generatePass(Length)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(myPass[0:Length])
}
