package main

import (
	"fmt"
	"os"
	"reflect"
)

type a struct {
	X int
	Y float64
	Z string
}

type b struct {
	F int
	G int
	H string
	I float64
}

func main() {
	x := 100
	xRefl := reflect.ValueOf(&x).Elem()
	xType := xRefl.Type()
	fmt.Printf("The type of x is %s.\n", xType)

	fmt.Println("The TypeOv value from x is", reflect.TypeOf(x))

	A := a{
		X: 100,
		Y: 300.12,
		Z: "Struct a",
	}
	B := b{
		F: 1,
		G: 2,
		H: "Struct b",
		I: -1.2,
	}
	var r reflect.Value
	args := os.Args[1:]
	if len(args) == 0 {
		r = reflect.ValueOf(&A).Elem()
	} else {
		r = reflect.ValueOf(&B).Elem()
	}

	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are:\n", r.NumField(), iType)

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("Field name: %s ", iType.Field(i).Name)
		fmt.Printf("with type: %s ", r.Field(i).Type())
		fmt.Printf("and value %v\n", r.Field(i).Interface())
	}
}
