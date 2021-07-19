package week12

import (
	. "algorithm-go/utils"
)

/*

单调递增的[]nums   (栈底)1, 2, 3, 5(栈顶)
s := stack<int>
for _, n := range nums{
	while len(s) != 0 && s.top() >= n{
		s.pop()
	}
	s.push(n)
}

1. 栈顶元素 -> 右边第一个比自己小的数字
2. 栈顶元素 -> 从自己往左第一个比自己小的数字
3. 栈顶元素 -> 包含自己，且自己是最大的一个子数组 (r-l)

4. 新数字 -> 从自己往左，第一个比自己大的数
5. 新数字 -> 左边最接近自己且比自己大的数
6. 新数字 -> 从自己往左，第一个比自己大的数

单调递减的[]nums   (栈底)5,3,2,1(栈顶)
s := stack<int>
for _, n := range nums{
	while len(s) != 0 && s.top() <= n{
		s.pop()
	}
	s.push(i)
}

解决的问题：
1. 左右第一个**
2. 找一个sub_array，自己最大或最小
*/

// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
func largestRectangleArea(heights []int) int {
	newHeight := make([]int, len(heights)+2)
	copy(newHeight[1:], heights)
	stack := make(IntStack, 0)
	ans := 0
	for i := 1; i < len(newHeight); i++ {
		for newHeight[i] < newHeight[stack.Top()] {
			// 如果循环到的元素的高度小于栈顶元素的高度，则栈顶元素的那个矩形的面积可以计算了
			// 高度为栈顶元素
			height := newHeight[stack.Pop()]
			// 宽度为次高的元素(新栈顶)到当前遍历到的元素的前一个的宽度
			width := i - 1 - stack.Top()
			ans = MaxInt(ans, height*width)
		}
		stack.Push(i)
	}
	return ans
}
