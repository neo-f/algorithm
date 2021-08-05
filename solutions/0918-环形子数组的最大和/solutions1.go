package main

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	// 前缀和数组
	pre := make([]int, 2*n+1)
	for i := 1; i < len(pre); i++ {
		pre[i] = pre[i-1] + nums[(i-1)%n]
	}

	// 单调队列求 i1 > i2 下 pre[i1] - pre[i2]的最大值
	//for i1 := 1; i1 <len(pre); i1++ {
	//	for i2 := 0; i2 < i1;i2++{
	//		ans = max(ans, pre[i1]-pre[i2])
	//	}
	//}
	deque := make([]int, 1)
	ans := -100000000
	for i := 1; i < len(pre); i++ {
		// 1. 队头合法性, 长度不超过n, i - deque[0] <= n
		for len(deque) != 0 && i-deque[0] > n {
			deque = deque[1:]
		}
		// 2. 取队头为最优
		ans = max(ans, pre[i]-pre[deque[0]])
		// 3. 保证队尾单调  pre[deque[-1]] < pre[i]
		for len(deque) != 0 && pre[deque[len(deque)-1]] >= pre[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
