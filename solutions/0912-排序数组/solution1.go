package main

import (
	"math/rand"
	"time"
)

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// 快排
func quickSort(nums []int, l, r int) {
	if l == r {
		return
	}
	pivot := partition(nums, l, r)
	quickSort(nums, l, pivot)
	quickSort(nums, pivot+1, r)
}

func partition(nums []int, l, r int) int {
	rand.Seed(time.Now().UnixNano())
	pivot := int(rand.Float64()*float64(r-l+1)) + l
	pivotVal := nums[pivot]
	for l <= r {
		for nums[l] < pivotVal {
			l++
		}
		for nums[r] > pivotVal {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	return r
}
