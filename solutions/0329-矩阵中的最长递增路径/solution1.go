package main

func longestIncreasingPath(matrix [][]int) int {
	// 记录坐标可以走的最大步数
	xLen := len(matrix)
	yLen := len(matrix[0])
	ans := 0

	steps := make([][]int, xLen)
	// setup default steps
	for x := 0; x < xLen; x++ {
		line := make([]int, yLen)
		for y := 0; y < yLen; y++ {
			line[y] = -1
		}
		steps[x] = line
	}

	//           U, R, D, L
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{-1, 0, 1, 0}

	var howFar func(x, y int) int
	howFar = func(x, y int) int {
		if step := steps[x][y]; step != -1 {
			return step
		}
		steps[x][y] = 1
		for d := 0; d < 4; d++ {
			nx, ny := x+dx[d], y+dy[d]
			if nx < 0 || ny < 0 || nx >= xLen || ny >= yLen {
				continue
			}
			if matrix[nx][ny] <= matrix[x][y] {
				continue
			}
			steps[x][y] = max(steps[x][y], howFar(nx, ny)+1)
		}
		return steps[x][y]
	}

	for x := 0; x < xLen; x++ {
		for y := 0; y < yLen; y++ {
			ans = max(ans, howFar(x, y))
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
