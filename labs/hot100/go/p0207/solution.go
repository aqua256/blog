package p0207

func canFinish(numCourses int, prerequisites [][]int) bool {
	// Kahn 拓扑排序：反复取入度为 0 的课，能取完说明无环
	graph := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, p := range prerequisites {
		course, pre := p[0], p[1]
		graph[pre] = append(graph[pre], course)
		indeg[course]++
	}
	queue := []int{}
	for i, d := range indeg {
		if d == 0 {
			queue = append(queue, i)
		}
	}
	taken := 0
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		taken++
		for _, v := range graph[u] {
			indeg[v]--
			if indeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return taken == numCourses
}
