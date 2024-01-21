package sudoku

func Parser(arguments []string) [9][9]int {
	// var sl []string
	var line [9][9]int
	for j, arg := range arguments {
		for i, char := range arg {
			if char == '.' {
				line[j][i] = 0
			} else {
				line[j][i] = int(rune(char) - '0')
			}
		}
	}
	// for i, number := range sl {
	// 	if number == string('.') {
	// 		line[i] = 0
	// 	} else {
	// 		line[i] =
	// 	}
	// }
	return line
}

func PrintSudoku(arguments []string /*sudoku [9][9]int*/) {
	/*for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			line[j][i] = Parser(arguments)[j][i]
		}
	}
*/
	sudoku := Parser(arguments)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", sudoku[i][j])
		}
		fmt.Println()
	}
}
// go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e