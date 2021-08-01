package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
	"os"
)

type Entry struct {
	Number int
	Double int
	Square int
}

var Data []Entry
var tFile string

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)
	myT := template.Must(template.ParseGlob(tFile))
	myT.ExecuteTemplate(w, tFile, Data)
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Need Database File + Template File!")
		return
	}

	database := args[1]
	tFile = args[2]

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Emptying database table")
	_, err = db.Exec("DELETE FROM data")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Populating", database)
	stmt, _ := db.Prepare("INSERT INTO data(number, double, square) values(?,?,?)")
	for i := 20; i < 50; i++ {
		_, _ = stmt.Exec(i, 2*i, i*i)
	}
	fmt.Println("Populating finished!")

	rows, err := db.Query("SELECT * FROM data")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rows)

	var n, d, s int
	for rows.Next() {
		err = rows.Scan(&n, &d, &s)
		temp := Entry{
			Number: n,
			Double: d,
			Square: s,
		}
		Data = append(Data, temp)
	}

	http.HandleFunc("/", myHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
