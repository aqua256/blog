package p0240

func searchMatrix(matrix [][]int, target int) bool {
	// 从右上角出发：左边都更小、下边都更大，每步排除一行或一列
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		switch v := matrix[i][j]; {
		case v == target:
			return true
		case v > target:
			j-- // 这一列往下只会更大
		default:
			i++ // 这一行往左只会更小
		}
	}
	return false
}
