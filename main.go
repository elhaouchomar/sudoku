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

	// now i have line wich is the 9 by 9 array
	// if it is valid i start locking for the solution,
	// if it not i return error and stop

	if sudoku.IsValid(line) {
		// Attempt to solve the Sudoku
		if sudoku.solveSudoku(&line, sudoku.IsValid) {
			// Print the solved Sudoku
			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					fmt.Printf("%d ", line[i][j])
				}
				fmt.Println()
			}
		}
	} else {
		fmt.Println("Error")
	}
}
