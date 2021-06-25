package week12

//https://leetcode-cn.com/problems/sliding-window-maximum/

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	q := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		// 1. 先按照时间淘汰
		for len(q) != 0 && q[0] <= i-k {
			q = q[1:]
		}
		// 2. 按照值淘汰
		for len(q) != 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if idx := i - k + 1; idx >= 0 {
			ans[idx] = nums[q[0]]
		}
	}
	return ans
}
