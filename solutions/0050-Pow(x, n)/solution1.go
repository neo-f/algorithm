package week21

func myPow(x float64, n int) float64 {
	// 负指数的情况
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	v := myPow(x, n/2)
	// 奇数
	if n&1 == 1 {
		return v * v * x
	}
	// 偶数
	return v * v
}
