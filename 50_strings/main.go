package main

import (
	"fmt"
	"strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	upper := strings.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", strings.ToLower("Hello THERE"))
	f("%s\n", strings.Title("tHis wiLL be A title!"))
	f("EqualFold: %v\n", strings.EqualFold("Zhenya", "ZHENya"))
	f("EqualFold: %v\n", strings.EqualFold("Zhenya", "ZHENy"))

	f("prefix: %v\n", strings.HasPrefix("Zhenya", "Zh"))
	f("prefix: %v\n", strings.HasPrefix("Zhenya", "zh"))
	f("suffix: %v\n", strings.HasSuffix("Zhenya", "ya"))
	f("suffix: %v\n", strings.HasSuffix("Zhenya", "Ya"))

	f("Index: %v\n", strings.Index("Zhenya", "en"))
	f("Index: %v\n", strings.Index("Zhenya", "En"))
	f("Count: %v\n", strings.Count("Zhenya", "n"))
	f("Count: %v\n", strings.Count("Zhenya", "N"))
	f("Repeat: %s\n", strings.Repeat("ab", 5))

	f("TrimSpace: %s\n", strings.TrimSpace(" \tThis is a line \n"))
	f("Trim left: %s\n", strings.TrimLeft(" \tThis is a\t line \n", "\n\t"))
	f("Trim right: %s\n", strings.TrimRight(" \tThis is a\t line \n", "\n\t"))

	f("Compare: %v\n", strings.Compare("Zhenya", "ZHENYA"))
	f("Compare: %v\n", strings.Compare("Zhenya", "Zhenya"))
	f("Compare: %v\n", strings.Compare("ZHENYA", "ZhEnYa"))

	f("Fields: %v\n", strings.Fields("This is a string!"))
	f("Fields: %v\n", strings.Fields("Thisis\na\tstring!"))

	f("%s\n", strings.Split("abcd efg", " "))

	f("%s\n", strings.Replace("abcd efg", "", "_", -1))
	f("%s\n", strings.Replace("abcd efg", "", "_", 4))
	f("%s\n", strings.Replace("abcd efg", "", "_", 2))

	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", strings.Join(lines, "+++"))
	f("Split After: %s\n", strings.SplitAfter("123++432++", "++"))

	trimFunctions := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", strings.TrimFunc("123 abc ABC \t .", trimFunctions))
}
