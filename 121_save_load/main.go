package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type myElement struct {
	Name    string
	Surname string
	Id      string
}

var Datafile = "/tmp/datafile.gob"
var Data = make(map[string]myElement)

func save() error {
	fmt.Println("Saving", Datafile)
	err := os.Remove(Datafile)
	if err != nil {
		fmt.Println(err)
	}

	saveTo, err := os.Create(Datafile)
	if err != nil {
		fmt.Println("Cannot create", Datafile)
		return err
	}
	defer saveTo.Close()

	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(Data)
	if err != nil {
		fmt.Println("Cannot save to", Data)
		return err
	}
	return nil
}

func load() error {
	fmt.Println("Loading", Datafile)
	loadFrom, err := os.Open(Datafile)
	defer loadFrom.Close()
	if err != nil {
		fmt.Println("Empty key/value store!")
		return err
	}

	decoder := gob.NewDecoder(loadFrom)
	decoder.Decode(&Data)
	return nil
}
