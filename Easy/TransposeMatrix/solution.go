package TransposeMatrix

func TransposeMatrix(matrix [][]int) (result [][]int) {
	lenY := len(matrix)
	lenX := len(matrix[0])

	result = make([][]int, lenX)

	for col := 0; col < lenX; col++ {
		newRow := make([]int, lenY)
		for row := 0; row < lenY; row++ {
			newRow[row] = matrix[row][col]
		}
		result[col] = newRow
	}

	return result
}
