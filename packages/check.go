package sub_Sudoku2

/*func Check_Line(line [9][9]int , number int)bool{
	// check if the number is unique on the line 

	for i := 0 ; i < 9 ; i ++ {
		if line[number]
	}
}
*/
/*func Check_Colone(){
	// check if the number is unique in the colon 

}*/

func Check_Square() {
	// check if the number is unique on a 3 by 3 square
}


/*func Check_And_Fille(array [9][9]) {
	// add a number . if there is only one place for it 
}*/
func Check_Line(line [9]int , number int , column int ) bool {
	// check if the number is unique on the line 
	for i := column + 1 ; i < 9 ; i++ {
		if line[i] == number  {
			return false
		}
	}
	return true
}
func Check_Colone(column [9]int , number int , row int) bool {
	// check if the number is unique in the colon 
	for i := row + 1 ; i < 9 ; i++ {
		if column[i] == number {
			return false
		}
	}
	return true
}
func ScanColumn(sudoku [9][9]int, column int) [9]int{
	columns := [9]int{}
	for i := 0; i < 9; i++ {
		columns[i] = sudoku[i][column]
	}
	return columns
}
func IsValid(sudoku [9][9]int) bool {
	isValid := true
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			Number := sudoku[row][column]
			if Number != 0 {
				isValid = (Check_Line(sudoku[row],Number,column) && Check_Colone(ScanColumn(sudoku,column),Number,row))
				if isValid == false {
					return false
				}
			}
		}
	}
	return isValid
}
