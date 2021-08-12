package _151_翻转字符串里的单词

import "strings"

func reverseWords(s string) string {
	parsed := strings.Fields(s)
	for left, right := 0, len(parsed)-1; left < right; {
		parsed[left], parsed[right] = parsed[right], parsed[left]
		left++
		right--
	}
	return strings.Join(parsed, " ")
}
