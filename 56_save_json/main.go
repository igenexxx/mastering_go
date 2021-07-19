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

func saveToJSON(filename *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	myRecord := Record{
		Name:    "Zhenya",
		Surname: "Sikirzhitsky",
		Tel: []Telephone{
			{
				Mobile: true,
				Number: "12345-abc",
			},
			{
				Mobile: false,
				Number: "abc-1234",
			},
			{
				Mobile: true,
				Number: "777-77-77",
			},
		},
	}

	saveToJSON(os.Stdout, myRecord)
}
