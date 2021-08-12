package _014_最长公共前缀

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// 获取最短的一个字符串
	minOne := strs[0]
	for _, str := range strs {
		if len(str) < len(minOne) {
			minOne = str
		}
	}
	// 遍历最短的这个字符串，所以不必判断索引越界问题
	for idx := range minOne {
		for _, str := range strs {
			if str[idx] != minOne[idx] {
				return minOne[:idx]
			}
		}
	}
	return minOne
}
