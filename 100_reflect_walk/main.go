package main

import (
	"fmt"
	"github.com/mitchellh/reflectwalk"
	"reflect"
)

type Values struct {
	Extra map[string]string
}

type WalkMap struct {
	MapVal reflect.Value
	Keys   map[string]bool
	Values map[string]bool
}

func (w *WalkMap) Map(m reflect.Value) error {
	w.MapVal = m
	return nil
}

func (w *WalkMap) MapElem(m, k, v reflect.Value) error {
	if w.Keys == nil {
		w.Keys = make(map[string]bool)
		w.Values = make(map[string]bool)
	}

	w.Keys[k.Interface().(string)] = true
	w.Keys[v.Interface().(string)] = true
	return nil
}

func main() {
	w := new(WalkMap)

	type S struct {
		Map map[string]string
	}

	data := &S{Map: map[string]string{
		"V1": "v1v",
		"V2": "v2v",
		"V3": "v3v",
		"V4": "v4v",
	}}

	err := reflectwalk.Walk(data, w)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := w.MapVal

	fmt.Println("MapVal:", r)
	rType := r.Type()
	fmt.Printf("Type of r: %s\n", rType)

	for _, key := range r.MapKeys() {
		fmt.Println("key", key, "value", r.MapIndex(key))
	}
}
