package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	column = iota
	row
)

func testRowColumn(sudoku [][]int, rowColumn int) bool {
	for i := 0; i < 9; i++ {
		sum := 0
		for j := 0; j < 9; j++ {
			if rowColumn == column {
				sum += sudoku[i][j]
			} else {
				sum += sudoku[j][i]
			}
		}
		if sum != 45 {
			return false
		}
		sum = 0
	}
	return true
}

func importFile(file string) ([][]int, error) {
	var err error
	var mySlice = make([][]int, 0)

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		fields := strings.Fields(line)
		temp := make([]int, 0)
		for _, v := range fields {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			temp = append(temp, n)
		}
		if len(temp) != 0 {
			mySlice = append(mySlice, temp)
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if len(temp) != len(mySlice[0]) {
			return nil, errors.New("Wrong number of elements!")
		}
	}

	return mySlice, nil
}

func validPuzzle(s1 [][]int) bool {
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			iE1 := i * 3
			jE1 := j * 3
			mySlice := make([]int, 9)
			for k := 0; k <= 2; k++ {
				for m := 0; m <= 2; m++ {
					bigI := iE1 + k
					bigJ := jE1 + m
					val := s1[bigI][bigJ]
					if val > 0 && val < 10 {
						if mySlice[val-1] == 1 {
							fmt.Println("Appear 2 times:", val)
							return false
						} else {
							mySlice[val-1] = 1
						}
					} else {
						fmt.Println("Invalid value:", val)
						return false
					}
				}
			}
		}
	}

	// Testing columns
	columnTestResult := testRowColumn(s1, column)
	if !columnTestResult {
		return false
	}

	rowTestResult := testRowColumn(s1, row)
	if !rowTestResult {
		return false
	}

	return true
}
