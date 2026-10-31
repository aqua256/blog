package p0054

func spiralOrder(matrix [][]int) []int {
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	ans := make([]int, 0, len(matrix)*len(matrix[0]))
	for top <= bottom && left <= right {
		for j := left; j <= right; j++ {
			ans = append(ans, matrix[top][j])
		}
		for i := top + 1; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		if top < bottom && left < right { // 只剩一行或一列时，不能再往回走
			for j := right - 1; j > left; j-- {
				ans = append(ans, matrix[bottom][j])
			}
			for i := bottom; i > top; i-- {
				ans = append(ans, matrix[i][left])
			}
		}
		top, bottom, left, right = top+1, bottom-1, left+1, right-1
	}
	return ans
}
