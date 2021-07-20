package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	// 第K大的值 -> 递减数组 nums[K-1]
	//			-> 递增数组 nums[N - 1 - (K - 1)]
	//			-> 递增数组 nums[N-K]的值
	return qsort(nums, 0, len(nums)-1, len(nums)-k)
}

func qsort(nums []int, left, right int, k int) int {
	if left == right {
		return nums[left]
	}
	pivot := partition(nums, left, right)
	if pivot >= k {
		// 分出来的中心点pivot坐标比K大，说明左边多，右边少，只需要继续处理左边部分
		return qsort(nums, left, pivot, k)
	}
	// 分出来的中心点pivot坐标比K小，说明左边少，右边多，只需要继续处理右边部分
	return qsort(nums, pivot+1, right, k)
}

func partition(nums []int, left, right int) int {
	pivot := left + 1
	pivotVal := nums[pivot]
	for left <= right {
		for nums[left] < pivotVal {
			left++
		}
		for nums[right] > pivotVal {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return right
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}

	fmt.Println(findKthLargest(nums, 2))
}
