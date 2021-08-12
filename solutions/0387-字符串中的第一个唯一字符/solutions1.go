package _387_字符串中的第一个唯一字符

func firstUniqChar(s string) int {
	// 哈希计数
	var count [26]int
	for _, w := range s {
		count[w-'a']++
	}
	// 碰到第一个计数为1的返回
	for idx, w := range s {
		if count[w-'a'] == 1 {
			return idx
		}
	}
	return -1
}
