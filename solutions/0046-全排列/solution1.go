package week21

// https://leetcode-cn.com/problems/permutations/
func permute(nums []int) [][]int {
	var ans [][]int
	var pmt []int
	used := make([]bool, len(nums))

	var find func(int)
	find = func(idx int) {
		// 终止条件
		if idx == len(nums) {
			ans = append(ans, append([]int(nil), pmt...))
			return
		}
		// 遍历数组
		for i := range nums {
			// 如果该值已经被使用了，do nothing
			// 如果没有被使用,
			// 1. pmt中加入该值，
			// 2. 标记该值已经被使用过了
			// 3. 递归下个索引
			// 4. 恢复临时变量
			if !used[i] {
				used[i] = true
				pmt = append(pmt, nums[i])
				find(idx + 1)
				used[i] = false
				pmt = pmt[:len(pmt)-1]
			}
		}
	}
	find(0)
	return ans
}
