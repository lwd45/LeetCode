package mid

func rotate(matrix [][]int) {
	N := len(matrix)
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
	for i, j := 0, N-1; i < j; i, j = i+1, j-1 {
		for k := 0; k < N; k++ {
			temp := matrix[k][i]
			matrix[k][i] = matrix[k][j]
			matrix[k][j] = temp
		}
	}
}
