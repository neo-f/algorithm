package main

func shipWithinDays(weights []int, days int) int {
	// 能否用weight装载力的船实现？
	canShipWithin := func(weight int) bool {
		// 临时累加值
		sum := 0
		// 趟数
		count := 1
		for _, w := range weights {
			sum += w
			if sum > weight {
				sum = w
				count++
				if count > days {
					return false
				}
			}
		}
		return count <= days
	}

	// 二分查找最小满足条件的值
	binarySearch := func(left, right int) int {
		for left < right {
			mid := (left + right) >> 1
			if canShipWithin(mid) {
				right = mid
			} else {
				left = mid + 1
			}
		}
		return left
	}

	// 最小条件为货物的最大重量
	// 最大条件为货物总重量
	max, sum := 0, 0
	for _, w := range weights {
		if w > max {
			max = w
		}
		sum += w
	}
	return binarySearch(max, sum)
}
