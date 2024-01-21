package main

import (
	"fmt"
	"os"

	sudoku "sudoku/packages"
)

func main() {
	var line [9][9]int
	arguments := os.Args[1:]
	// )
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			line[i][j] = sudoku.Parser(arguments)[i][j]
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", line[i][j])
		}
		fmt.Println()
	}
}
