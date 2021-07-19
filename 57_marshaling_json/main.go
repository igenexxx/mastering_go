package main

import (
	"encoding/json"
	"fmt"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	IsMobile bool
	Number   string
}

func main() {
	myRecord := Record{
		Name:    "Zhenya",
		Surname: "Sikirzhitsky",
		Tel: []Telephone{
			{
				IsMobile: true,
				Number:   "123-abc",
			},
			{
				IsMobile: false,
				Number:   "abc-123",
			},
			{
				IsMobile: true,
				Number:   "123-456",
			},
		},
	}

	rec, err := json.Marshal(&myRecord)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(rec))

	var unRec Record
	err = json.Unmarshal(rec, &unRec)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(unRec)
}
