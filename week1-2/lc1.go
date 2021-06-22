package week12

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{v, i}
		}
		m[num] = i
	}
	return nil
}
