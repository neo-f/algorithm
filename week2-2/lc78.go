package week21

// https://leetcode-cn.com/problems/subsets/

// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

func subsets(nums []int) [][]int {
	var ans [][]int
	var shared []int
	var finder func(int)
	finder = func(idx int) {
		if idx == len(nums) {
			v := make([]int, len(shared))
			copy(v, shared)
			ans = append(ans, v)
			return
		}
		finder(idx + 1)
		shared = append(shared, nums[idx])
		finder(idx + 1)
		shared = shared[:len(shared)-1]
	}
	finder(0)
	return ans
}
