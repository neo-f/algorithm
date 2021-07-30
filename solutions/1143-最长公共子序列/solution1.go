package main

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	// 方便处理边界
	text1 = " " + text1
	text2 = " " + text2

	status := make([][]int, m+1)
	for i := range status {
		status[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i] == text2[j] {
				status[i][j] = status[i-1][j-1]
			} else {
				status[i][j] = max(status[i-1][j], status[i][j-1])
			}
		}
	}
	return status[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
