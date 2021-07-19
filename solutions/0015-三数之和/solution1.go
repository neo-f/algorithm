package week12

import "sort"

func threeSum(nums []int) [][]int {
	var results [][]int
	sort.Ints(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		result := getTwoSum(nums[i+1:], -nums[i])
		for _, pair := range result {
			results = append(results, []int{nums[i], pair[0], pair[1]})
		}
	}
	return results
}

func getTwoSum(nums []int, target int) [][2]int {
	var ret [][2]int
	i, j := 0, len(nums)-1
	for i < j {
		if i > 0 && nums[i] == nums[i-1] {
			i++
			continue
		}
		sum := nums[i] + nums[j]
		switch {
		case sum > target:
			j--
		case sum < target:
			i++
		default:
			ret = append(ret, [2]int{nums[i], nums[j]})
			i++
			j--
		}
	}
	return ret
}
