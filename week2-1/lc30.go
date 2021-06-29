package week21

func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	wordsLen := wordLen * len(words)

	wordsCount := make(map[string]int, 0)
	ans := make([]int, 0)
	for _, s := range words {
		wordsCount[s]++
	}
	for i := 0; i+wordsLen <= len(s); i++ {
		count := make(map[string]int)
		for j := 0; j < wordsLen; j += wordLen {
			count[s[i+j:i+j+wordLen]]++
		}

		if mapIsEqual(count, wordsCount) {
			ans = append(ans, i)
		}
	}
	return ans
}

func mapIsEqual(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; !ok || v2 != v1 {
			return false
		}
	}
	for k2, v2 := range m2 {
		if v1, ok := m1[k2]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}
