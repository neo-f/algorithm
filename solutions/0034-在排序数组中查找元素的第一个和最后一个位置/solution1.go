package main

func searchRange(nums []int, target int) []int {
	var l, r int
	// 1. 寻找第一个>= target的值
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	l = left

	// 2. 寻找最后一个<= target的值
	left, right = -1, len(nums)-1
	for left < right {
		mid := (left + right + 1) >> 1
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	r = left

	if l > r {
		return []int{-1, -1}
	}
	return []int{l, r}
}
