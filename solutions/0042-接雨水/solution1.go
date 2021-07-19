package week12

import (
	. "algorithm-go/utils"
)

func trap(height []int) int {
	var ans int
	stack := make(IntStack, 0)
	for i := range height {
		// 遇到的柱子更高时计算前面的值
		for !stack.IsEmpty() && height[i] > height[stack.Top()] {
			// 低洼的值
			bottom := stack.Pop()
			// 防止越界
			if stack.IsEmpty() {
				continue
			}
			// 左右边界
			left, right := stack.Top(), i
			// 高度
			h := MinInt(height[left], height[right]) - height[bottom]
			// 宽度
			w := i - 1 - left
			ans += h * w
		}
		stack.Push(i)
	}

	return ans
}
