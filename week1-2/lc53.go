package week12

import "math"

//https://leetcode-cn.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	acc := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		acc[i] = acc[i-1] + nums[i-1]
	}
	ans := math.MinInt64
	preMin := acc[0]
	for r := 1; r <= len(nums); r++ {
		// 最大值 -> acc[r] - acc[l] 最大 -> acc[l] 最小
		ans = maxInt(ans, acc[r]-preMin)
		preMin = minInt(preMin, acc[r])
	}
	return ans
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func minInt(i, j int) int {
	if i > j {
		return j
	}
	return i
}
