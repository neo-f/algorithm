package main

import "math"

func coinChange(coins []int, amount int) int {
	status := make([]int, amount+1)
	for i := range status {
		status[i] = -1
	}

	var calc func(amt int) int
	calc = func(amt int) int {
		if amt == 0 {
			return 0
		}
		if amt < 0 {
			return math.MaxInt32
		}
		if status[amt] != -1 {
			return status[amt]
		}
		status[amt] = math.MaxInt32
		for _, coin := range coins {
			status[amt] = min(status[amt], calc(amt-coin)+1)
		}
		return status[amt]
	}
	v := calc(amount)
	if v >= math.MaxInt32 {
		return -1
	}
	return v
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
