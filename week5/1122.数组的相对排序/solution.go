package main

// 数组解法
func relativeSortArray(arr1 []int, arr2 []int) []int {
	var max int
	for _, v := range arr1 {
		if v > max {
			max = v
		}
	}

	counter := make([]int, max+1)
	for _, v := range arr1 {
		counter[v]++
	}

	var ans []int
	for _, v := range arr2 {
		if count := counter[v]; count > 0 {
			for i := 0; i < count; i++ {
				ans = append(ans, v)
				counter[v]--
			}
		}
	}
	for v, count := range counter {
		if count > 0 {
			for i := 0; i < count; i++ {
				ans = append(ans, v)
			}
		}
	}
	return ans
}
