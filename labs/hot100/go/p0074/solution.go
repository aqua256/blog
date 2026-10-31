package p0074

// 行首 > 上一行行尾，所以整个矩阵按行展开就是一个有序数组：
// 下标 k 对应 matrix[k/n][k%n]，在 [0, m*n) 上做 lowerBound。
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	lo, hi := 0, m*n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if matrix[mid/n][mid%n] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo < m*n && matrix[lo/n][lo%n] == target
}
