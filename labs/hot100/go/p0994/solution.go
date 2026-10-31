package p0994

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var queue [][2]int
	fresh := 0
	for i := range m {
		for j := range n {
			switch grid[i][j] {
			case 1:
				fresh++
			case 2:
				queue = append(queue, [2]int{i, j}) // 所有烂橘子同时作为起点
			}
		}
	}
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	minutes := 0
	for len(queue) > 0 && fresh > 0 {
		minutes++
		next := [][2]int{}
		for _, p := range queue {
			for _, d := range dirs {
				x, y := p[0]+d[0], p[1]+d[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					grid[x][y] = 2
					fresh--
					next = append(next, [2]int{x, y})
				}
			}
		}
		queue = next
	}
	if fresh > 0 {
		return -1
	}
	return minutes
}
