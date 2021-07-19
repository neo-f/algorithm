package main

import (
	"math/rand"
	"time"
)

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	// mergeSort(nums, 0, len(nums)-1)
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

// 归并排序

func mergeSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) >> 1
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int) {
	var tmp []int
	i, j := left, mid+1
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			tmp = append(tmp, nums[i])
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for i <= mid {
		tmp = append(tmp, nums[i])
		i++
	}
	for j <= right {
		tmp = append(tmp, nums[j])
		j++
	}
	copy(nums[left:right+1], tmp)
}
