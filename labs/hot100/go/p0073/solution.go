package p0073

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	// 用第 0 行/第 0 列当标记位，它们自己是否要清零先单独记下
	row0, col0 := false, false
	for j := range n {
		row0 = row0 || matrix[0][j] == 0
	}
	for i := range m {
		col0 = col0 || matrix[i][0] == 0
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row0 {
		for j := range n {
			matrix[0][j] = 0
		}
	}
	if col0 {
		for i := range m {
			matrix[i][0] = 0
		}
	}
}
