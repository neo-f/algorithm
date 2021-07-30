package main

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	maxX := len(obstacleGrid)
	maxY := len(obstacleGrid[0])

	status := make([][]int, maxX)
	for i := range status {
		status[i] = make([]int, maxY)
		for j := range status[i] {
			if obstacleGrid[i][j] == 1 {
				// 障碍物
				status[i][j] = 0
			} else if i == 0 && j == 0 {
				// 起点
				status[i][j] = 1
			} else if i == 0 {
				// 顶边
				status[i][j] = status[i][j-1]
			} else if j == 0 {
				// 底边
				status[i][j] = status[i-1][j]
			} else {
				status[i][j] = status[i-1][j] + status[i][j-1]
			}
		}
	}
	return status[maxX-1][maxY-1]
}
