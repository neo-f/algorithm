package week12

func twoSum2(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	// 数组是递增的，所以numbers[i] 永远小于等于 numbers[j]
	for i < j {
		sum := numbers[i] + numbers[j]
		switch {
		// 值太大，需要变小，所以右侧指针向左移动
		case sum > target:
			j--
		// 值太小，需要变大，所以左侧指针向右移动
		case sum < target:
			i++
		case sum == target:
			return []int{i + 1, j + 1}
		}
	}
	return nil
}
