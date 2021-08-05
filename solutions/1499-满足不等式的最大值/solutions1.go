package _499_满足不等式的最大值

import "math"

type intQueue []int

func (q intQueue) Empty() bool {
	return len(q) == 0
}

func (q intQueue) Front() int {
	return q[0]
}
func (q *intQueue) PopFront() {
	*q = (*q)[1:]
}
func (q intQueue) Back() int {
	return q[len(q)-1]
}
func (q *intQueue) PopBack() {
	*q = (*q)[:len(*q)-1]
}

func (q *intQueue) PushBack(i int) {
	*q = append(*q, i)
}

func findMaxValueOfEquation(points [][]int, k int) int {
	deque := make(intQueue, 0)
	ans := math.MinInt64
	for i := range points {

		// j 的上界: i - 1  (单调递增)
		// j 的下界满足 xi - xj < k  =>   xj >= xi - k (单调递增)
		// -> 滑动窗口

		// y[i] - x[i] 的单调递增队列
		// 满足 j1 < j2, j1 比 j2 优的条件: y[j1] - x[j1] > y[j2] - x[j2]

		// 1. 队头合法性 (不满足:34的pop掉)
		for !deque.Empty() && points[deque.Front()][0] < points[i][0] - k {
			deque.PopFront()
		}
		// 2. 队头最优解
		if !deque.Empty(){
			ans = max(ans, points[i][1] + points[i][0] + points[deque.Front()][1] - points[deque.Front()][0])
		}
		// 3. 维护队尾单调递增(不满足:38的pop掉)
		for !deque.Empty() && points[deque.Back()][1] - points[deque.Back()][0] <= points[i][1] - points[i][0]{
			deque.PopBack()
		}
		deque.PushBack(i)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
