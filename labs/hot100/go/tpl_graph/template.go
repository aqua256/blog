package tpl_graph

var dirs = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

// 网格 DFS：越界 / 不合法 / 已访问 三连判断后直接 return，访问即标记
func floodFill(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	for _, d := range dirs {
		floodFill(grid, i+d[0], j+d[1])
	}
}

// 多源 BFS 按层扩散：返回每格到最近源点的步数（不可达为 -1）
func bfsLevels(grid [][]int, sources [][2]int) [][]int {
	m, n := len(grid), len(grid[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	queue := sources
	for _, s := range sources {
		dist[s[0]][s[1]] = 0
	}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			x, y := p[0]+d[0], p[1]+d[1]
			if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] != 0 && dist[x][y] == -1 {
				dist[x][y] = dist[p[0]][p[1]] + 1 // 入队时标记，避免重复入队
				queue = append(queue, [2]int{x, y})
			}
		}
	}
	return dist
}

// Kahn 拓扑排序：edges[i] = {from, to}；有环时返回的序列长度 < n
func topoSort(n int, edges [][2]int) []int {
	graph := make([][]int, n)
	indeg := make([]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		indeg[e[1]]++
	}
	order := []int{}
	for i := range n {
		if indeg[i] == 0 {
			order = append(order, i)
		}
	}
	for i := 0; i < len(order); i++ { // order 兼作队列
		for _, v := range graph[order[i]] {
			if indeg[v]--; indeg[v] == 0 {
				order = append(order, v)
			}
		}
	}
	return order
}
