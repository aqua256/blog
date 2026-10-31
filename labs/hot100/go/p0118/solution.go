package p0118

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := range numRows {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = res[i-1][j-1] + res[i-1][j] // 左上 + 正上
		}
		res[i] = row
	}
	return res
}
