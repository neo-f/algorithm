package week1_1

// 560. 和为K的子数组 https://leetcode-cn.com/problems/subarray-sum-equals-k/
func subarraySum(nums []int, k int) int {
	// 前缀和计数，K=>前缀和，V=>次数
	s := map[int]int{0: 1}
	count := 0
	pre := 0
	for i := 0; i < len(nums); i++ {
		// 当前项的前缀和
		pre += nums[i]
		if c, ok := s[pre-k]; ok {
			count += c
		}
		s[pre]++
	}
	return count
}
