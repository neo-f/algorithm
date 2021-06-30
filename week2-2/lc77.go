package week21

// https://leetcode-cn.com/problems/combinations/

func combine(n int, k int) [][]int {
	var shared []int
	var ans [][]int

	var finder func(int)
	finder = func(num int) {
		// 优化
		// 1. shared大小超过了要选的数量
		// 2. 剩下的全选也不够k个
		if len(shared) > k || (n-num)+len(shared)+1 < k {
			return
		}
		// 终止条件
		if num == n+1 {
			ans = append(ans, append([]int{}, shared...))
			return
		}
		finder(num + 1)
		shared = append(shared, num)
		finder(num + 1)
		shared = shared[:len(shared)-1]
	}
	finder(1)
	return ans
}
