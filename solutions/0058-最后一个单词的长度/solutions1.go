package _058_最后一个单词的长度

func lengthOfLastWord(s string) int {
	started := false
	var count int
	for i := len(s) - 1; i >= 0; i-- {
		// 如果是空格
		if s[i] == ' ' && started {
			if started {
				// 且已经开始数字母, 说明单词已经结束
				return count
			} else {
				// 且未开始数字母, 忽略之
				continue
			}
		} else {
			// 如果不是空格，则计数，并标记已经开始数字母
			started = true
			count++
		}

	}
	return count
}
