package week21

//https://leetcode-cn.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	set := make(map[[26]int][]string)
	for _, s := range strs {
		set[groupAnagramsCalcHash(s)] = append(set[groupAnagramsCalcHash(s)], s)
	}
	ans := make([][]string, 0, len(set))
	for _, strings := range set {
		ans = append(ans, strings)
	}
	return ans
}

func groupAnagramsCalcHash(s string) [26]int {
	// Go语言中只要是可以比较大小的值都可以用作map的KeyType, 参考https://golang.org/ref/spec#Comparison_operators
	var ans [26]int
	for _, r := range s {
		ans[r-'a']++
	}
	return ans
}
