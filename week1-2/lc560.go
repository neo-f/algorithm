package week12

//https://leetcode-cn.com/problems/subarray-sum-equals-k/
func subarraySum(nums []int, k int) int {
	count := map[int]int{0: 1}
	pre, ans := 0, 0
	for i := range nums {
		pre += nums[i]
		//pre[j] - pre[i] = k
		//pre[i] = pre[j] - k
		if c, ok := count[pre-k]; ok {
			ans += c
		}
		count[pre]++
	}
	return ans
}
