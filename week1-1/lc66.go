package week1_1

// 66. 加一 https://leetcode-cn.com/problems/plus-one/
// 输入：digits = [1,2,3]
// 输出：[1,2,4]

func plusOne(digits []int) []int {
	ptr := len(digits) - 1
	plus := true
	for ptr >= 0 {
		if digits[ptr] == 9 && plus {
			// 如果是9且需要进位, 则数字改为0, 保持进位
			digits[ptr] = 0
		} else if plus {
			// 如果不是9且需要进位, 则数字+1, 不再进位
			digits[ptr]++
			plus = false
		} else {
			// 如果不是9且不需要进位, 不动, 且不再进位
			plus = false
		}
		ptr--
	}
	// 如果结果仍然需要进位，则头补1
	if plus {
		digits = append([]int{1}, digits...)
	}
	return digits
}
