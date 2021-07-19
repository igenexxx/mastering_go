package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}

	filename := args[1]

	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	var parsedData map[string]interface{}
	json.Unmarshal(fileData, &parsedData)

	for key, value := range parsedData {
		fmt.Printf("Key: %v\nValue: %v\n---", key, value)
	}
}
