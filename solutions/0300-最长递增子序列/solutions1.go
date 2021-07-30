package main

func lengthOfLIS(nums []int) int {
	store := make([]int, len(nums)+1)
	for i := range store {
		store[i] = 1
	}
	for r := 1; r < len(nums); r++ {
		for l := 0; l < r; l++ {
			if nums[l] < nums[r] {
				store[r] = max(store[l]+1, store[r])
			}
		}
	}
	ans := -10000
	for _, n := range store {
		if n > ans {
			ans = n
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
