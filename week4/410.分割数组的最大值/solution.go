package solution

// https://leetcode-cn.com/problems/split-array-largest-sum/

func splitArray(nums []int, m int) int {

	// 判定数组是否可以被分为小于m个，且每段和小于limit
	canSplitWithin := func(limit int) bool {
		sum, count := 0, 1
		for _, n := range nums {
			if sum+n <= limit {
				sum += n
			} else {
				sum = n
				count++
				if count > m {
					return false
				}
			}
		}
		return true
	}

	// 二分找Limit
	// 1. 找上限下限，上限为所有数字之和，下限为最大的那个数
	var left, right int
	for _, n := range nums {
		// 数字都为非负数
		if n > left {
			left = n
		}
		right += n
	}
	// 2. 找limit
	for left < right {
		mid := (left + right) >> 1
		// 能满足条件，向左
		if canSplitWithin(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
