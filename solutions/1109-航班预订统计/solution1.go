package week12

// 1109. 航班预订统计 https://leetcode-cn.com/problems/corporate-flight-bookings/
func corpFlightBookings(bookings [][]int, n int) []int {
	// 差分数组分别在头尾各多开一个空间
	delta := make([]int, n+2)
	for _, book := range bookings {
		delta[book[0]] += book[2]
		delta[book[1]+1] -= book[2]
	}
	// 累加数组也要多开一个空间
	acc := make([]int, n+1)
	for i := 1; i <= n; i++ {
		acc[i] = acc[i-1] + delta[i]
	}
	return acc[1:]
}
