package _771_宝石与石头

func numJewelsInStones(jewels string, stones string) int {
	// 计数数组，空间为 [A,B,C...a,b,c...z]
	count := make([]int, 'z'-'A'+1)
	for _, stone := range stones {
		count[stone-'A']++
	}
	var ans int
	for _, jewel := range jewels {
		ans += count[jewel-'A']
	}
	return ans
}
