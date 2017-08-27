package matrixElementsSum

func matrixElementsSum(matrix [][]int) int {
	// final price
	price := 0
	numOfRows := len(matrix)
	numOfColumns := len(matrix[0])
	// The question was not clear enough, they wanted everything below a 0 room
	// to be discarded.
	// make an slice of 0s, if we found a haunted room, mark the column haunted
	skipping := make([]int, numOfColumns, numOfColumns)
	for i := 0; i < numOfRows; i++ {
		for j := 0; j < numOfColumns; j++ {
			// if this column is marked, skip
			if skipping[j] > 0 {
				continue
			}
			// if found 0 room, mark current column as skipping
			if matrix[i][j] == 0 {
				skipping[j]++
			}

			// if not skipping, add price
			price += matrix[i][j]
		}
	}
	return price
}
