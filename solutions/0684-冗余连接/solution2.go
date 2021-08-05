package _684_冗余连接

// dfs
func findRedundantConnection2(input [][]int) []int {
	// 练习dfs

	// 获取有多少个节点
	n := 0
	for i := range input {
		x, y := input[i][0], input[i][1]
		if x > n {
			n = x
		}
		if y > n {
			n = y
		}
	}

	// n从1开始, 所以会访问edges[n], 所以开辟的空间+1
	// 初始化出边数组
	edges := make([][]int, n+1)
	visited := make([]bool, n+1)
	hasCycle := false

	var dfs func(x, parent int)
	dfs = func(x, parent int) {
		visited[x] = true
		for _, y := range edges[x] {
			// 无向图，父节点来的不管，保持单向
			if y == parent {
				continue
			}
			//
			if visited[y] {
				hasCycle = true
				return
			}
			// 所有的y都来自x（出边数组定义）
			dfs(y, x)
		}
		visited[x] = false
	}

	for i := range input {
		x, y := input[i][0], input[i][1]
		// 无向
		edges[x] = append(edges[x], y)
		edges[y] = append(edges[y], x)
		// 此处x/y应该都可以
		dfs(y, -1)
		if hasCycle {
			return input[i]
		}
	}
	return nil
}
