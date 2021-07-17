package main

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		lmid := (left + right) >> 1
		rmid := lmid + 1
		if nums[lmid] < nums[rmid] {
			// 不可能在lmid左边
			left = lmid + 1
		} else {
			right = rmid - 1
		}
	}
	return left
}
