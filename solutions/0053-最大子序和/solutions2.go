package week12

// 动态规划
func maxSubArray2(nums []int) int {
	f := make([]int, len(nums)+1)
	f[0] = nums[0]
	ans := f[0]
	for i := 1; i <= len(nums); i++ {
		f[i] = max(nums[i], nums[i]+f[i-1])
		ans = max(ans, f[i])
	}
	return ans
}
