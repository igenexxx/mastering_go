package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
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

func loadFromJson(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}

	filename := args[1]

	var myRecord Record
	err := loadFromJson(filename, &myRecord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("JSON:", myRecord)

	myRecord.Name = "Zhenya"

	xmlData, _ := xml.MarshalIndent(myRecord, "", "	")
	xmlData = []byte(xml.Header + string(xmlData))
	fmt.Println("\nxmlData:", string(xmlData))

	data := &Record{}
	err = xml.Unmarshal(xmlData, data)
	if err != nil {
		fmt.Println("Unmarshalling from XML", err)
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling to JSON", err)
		return
	}

	_ = json.Unmarshal(result, &myRecord)
	fmt.Println("\nJSON:", myRecord)
}
