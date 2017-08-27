package matrixElementsSum

import "testing"

var matrix = [][]int{
	[]int{0, 1, 1, 2},
	[]int{0, 5, 0, 0},
	[]int{2, 0, 3, 3},
}

func TestMatrixElementsSum(t *testing.T) {
	if matrixElementsSum(matrix) != 9 {
		t.Error("no")
	}
}
