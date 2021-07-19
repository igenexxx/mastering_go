package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func loadFromJson(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	decodeJson := json.NewDecoder(in)
	err = decodeJson.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a filename")
		return
	}

	filename := args[1]

	var myRecord Record
	err := loadFromJson(filename, &myRecord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myRecord)
}
