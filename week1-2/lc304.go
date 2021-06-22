package week12

// 304. 二维区域和检索 - 矩阵不可变 https://leetcode-cn.com/problems/range-sum-query-2d-immutable/

type NumMatrix struct {
	accMatrix [][]int
}

func NewNumMatrix(matrix [][]int) NumMatrix {
	// 计算二维前缀和
	accMatrix := make([][]int, len(matrix))
	for i := range matrix {
		accMatrix[i] = make([]int, len(matrix[i]))
		for j := range matrix[i] {
			accMatrix[i][j] = getAccSum(accMatrix, i-1, j) + getAccSum(accMatrix, i, j-1) - getAccSum(accMatrix, i-1, j-1) + matrix[i][j]
		}
	}
	return NumMatrix{accMatrix: accMatrix}
}

func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return getAccSum(m.accMatrix, row2, col2) - getAccSum(m.accMatrix, row2, col1-1) - getAccSum(m.accMatrix, row1-1, col2) + getAccSum(m.accMatrix, row1-1, col1-1)
}
func getAccSum(acc [][]int, i, j int) int {
	// i = 0 或者 j = 0 代表坐标是在边上
	if i >= 0 && j >= 0 {
		return acc[i][j]
	}
	return 0
}
