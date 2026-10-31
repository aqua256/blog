package p0048

import "slices"

func rotate(matrix [][]int) {
	n := len(matrix)
	// 顺时针 90° = 先沿主对角线转置，再每行左右翻转
	for i := range n {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for _, row := range matrix {
		slices.Reverse(row)
	}
}
