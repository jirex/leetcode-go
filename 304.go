type NumMatrix struct {
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var rows, cols int

	if len(matrix) == 0 {
		rows, cols = 1, 1
	} else {
		rows, cols = len(matrix)+1, len(matrix[0])+1
	}

	preSum := make([][]int, rows)

	for i := range preSum {
		preSum[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 || j == 0 {
				preSum[i][j] = 0
			} else {
				preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + matrix[i-1][j-1]
			}
		}
	}

	return NumMatrix{preSum: preSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] + this.preSum[row1][col1] - this.preSum[row2+1][col1] - this.preSum[row1][col2+1]
}
