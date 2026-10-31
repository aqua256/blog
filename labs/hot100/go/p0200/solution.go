package p0200

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var sink func(i, j int)
	sink = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '0' // 淹掉，相当于标记已访问
		sink(i+1, j)
		sink(i-1, j)
		sink(i, j+1)
		sink(i, j-1)
	}
	count := 0
	for i := range m {
		for j := range n {
			if grid[i][j] == '1' {
				count++
				sink(i, j)
			}
		}
	}
	return count
}
