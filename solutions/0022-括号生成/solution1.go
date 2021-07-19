package week21

import "fmt"

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	// 分组方法(a)b
	// generateSub(a), generateSub(b)
	var ans []string
	// 0的情况已经排除，从1开始遍历
	// a始终少一个初始化的括号，所以 a -> k-1
	// b的推导流程
	//-> a + b + 1 = n       a的括号数+b的括号数+初始化的1 = 总括号数量
	//-> k - 1 + b + 1 = n   代入 a = k-1
	//-> k + b = n
	//-> b = n - k
	for k := 1; k <= n; k++ {
		a := generateParenthesis(k - 1)
		b := generateParenthesis(n - k)
		for _, va := range a {
			for _, vb := range b {
				ans = append(ans, fmt.Sprintf("(%s)%s", va, vb))
			}
		}
	}
	return ans
}
