package main

import "math"

// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
// 你可以认为每种硬币的数量是无限的。

func coinChange(coins []int, amount int) int {
	status := make([]int, amount+1)
	status[0] = 0
	for i := 1; i <= amount; i++ {
		status[i] = math.MaxInt32
		for _, coin := range coins {
			if i-coin >= 0 {
				status[i] = min(status[i], status[i-coin]+1)
			}
		}
	}
	if status[amount] == math.MaxInt32 {
		return -1
	}
	return status[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
