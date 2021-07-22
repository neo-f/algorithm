package main

func minEatingSpeed(piles []int, h int) int {

	// 子问题分解，能否用k的速度完成任务？
	canEatIn := func(k int) bool {
		hours := 0
		for _, pile := range piles {
			// 求出 pile / k 的向上取整的值
			s := pile / k
			if pile%k != 0 {
				s += 1
			}
			hours += s
		}
		return hours <= h
	}

	// 二分求解，最慢是1小时吃1根，最快是1小时能吃完任意的一堆(最多的那堆一小时吃完)
	left, right := 1, 0
	for _, p := range piles {
		if p > right {
			right = p
		}
	}

	for left < right {
		mid := (left + right) >> 1
		if canEatIn(mid) {
			// 如果能吃完，就还有向左移动区间的余地，并包含该中间值[left, mid]
			right = mid
		} else {
			// 否则需要向右区间，不包含该中间值 [mid + 1, right]
			left = mid + 1
		}
	}
	return left
}
