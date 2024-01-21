package sudoku



func Solve_Soduko(array [9][9]int, x int, y int) [9][9]int {
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			if array[i][j] == 0 {
				for solution := 1; solution <= 9; solution++ {
					array[i][j] = solution
					if IsValid(array) {
						Stro_Changes(i, j, solution)
						break
					} else {
						continue
					}
				}
			}
		}
	}
	return array 
}

////////////////////////////////////////////////
type Coordinate struct {
	Row, Col int
}
 
func Stro_Changes(i int, j int, solution int) map[Coordinate]int {

    changes := make(map[Coordinate]int)
    changes[Coordinate{i, j}] = solution

    return changes
}
//////////////////////////////////////////////

func Back_Track(){

}