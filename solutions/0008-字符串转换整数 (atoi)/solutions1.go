package _008_字符串转换整数__atoi_

import "math"

func myAtoi(s string) int {
	idx := 0
	// 1. 忽略prefix空格
	for idx < len(s) && s[idx] == ' ' {
		idx++
	}
	var val int
	sign := 1
	// 2. 判断符号
	if idx < len(s) && s[idx] == '+' {
		idx++
	} else if idx < len(s) && s[idx] == '-' {
		sign = -1
		idx++
	}
	for idx < len(s) && s[idx] >= '0' && s[idx] <= '9' {
		// 3. 判断越界 val * 10 + int(s[idx] - '0') > math.MaxInt32
		if val > (math.MaxInt32-int(s[idx])+'0')/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		val = val*10 + int(s[idx]-'0')
		idx++
	}
	return val * sign
}
