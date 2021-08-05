package _200_岛屿数量

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	disjointSet := make([]int, n*m+1)
	idx := func(i, j int) int {
		return i*len(grid[0]) + j
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			disjointSet[idx(i, j)] = idx(i, j)
		}
	}
	// 外部海洋
	disjointSet[m*n] = m * n

	var find func(x int) int
	find = func(x int) int {
		if disjointSet[x] == x {
			return x
		}
		disjointSet[x] = find(disjointSet[x])
		return disjointSet[x]
	}

	union := func(x, y int) {
		x, y = find(x), find(y)
		if x != y {
			disjointSet[x] = y
		}
	}

	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, -1, 0, 1}
	for i := range grid {
		for j := range grid[i] {
			// 如果是水，就连到外部的海洋
			if grid[i][j] == '0' {
				union(idx(i, j), m*n)
				continue
			}
			// 如果是陆地，遍历四个方向，并连接相邻的陆地
			for d := 0; d < 4; d++ {
				ni := i + dx[d]
				nj := j + dy[d]
				if ni >= 0 && nj >= 0 && ni < m && nj < n && grid[ni][nj] == '1' {
					union(idx(i, j), idx(ni, nj))
				}
			}
		}
	}

	// 将并查集中不同的数字的数量即集合的数量，由于水的部分都属于外部海洋，那么其他的不同的数字都是一个单独的岛屿
	s := make(map[int]struct{})
	for _, num := range disjointSet {
		s[find(num)] = struct{}{}
	}
	// 岛屿的数量就是所有集合的数量减去外部的海洋
	return len(s) - 1
}
