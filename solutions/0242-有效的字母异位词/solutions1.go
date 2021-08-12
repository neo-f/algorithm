package _242_有效的字母异位词

func isAnagram(s string, t string) bool {
	return counter(s) == counter(t)
}

func counter(s string) [26]int {
	var c [26]int
	for _, r := range s {
		c[r-'a']++
	}
	return c
}
