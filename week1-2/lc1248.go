package week12

// 1248. 统计「优美子数组」 https://leetcode-cn.com/problems/count-number-of-nice-subarrays/
func numberOfSubarrays(nums []int, k int) (ans int) {
	accNums := make([]int, len(nums)+1)
	// 计数数组，统计和为index出现的次数 (index为数值，value为次数)
	count := make([]int, len(accNums))
	count[0] = 1
	// nums下标的取值范围是0~n-1
	// accNums下标的取值范围是0~n, 因为比nums多一个首位0
	// 遍历从 i = 1 ~ n 可满足 i-1 = 0 ~ n-1
	for i := 1; i < len(accNums); i++ {
		accNums[i] = accNums[i-1] + nums[i-1]&1
		// accNums[i] 即0~i的数组里累加等于该值
		// 也就是0到i的数组里面有accNums[i]个奇数（偶数已经都变为0了）
		// 0~i的前缀和出现次数+1
		// count[累加值] += 1
		count[accNums[i]]++
	}
	// for right = 0 ~ n -1
	//		for left = 0 ~ right
	//			if nums[right] - nums[left] == k:
	//          	ans += 1
	//  转化流程:
	//  固定外层变量right
	//  1. 对于每个right 有多少个 left(0~right) 满足accNums[right] - accNums[left] == k
	//  2. 对于每个right 有多少个 left(0~right) 满足accNums[left] == accNums[right] -  k
	//  3. 对于每个right 有多少个 accNums[left] == accNums[right] -  k
	//  4. 累加每个right 对应(accNums[right] -  k) 的数量

	for i := range accNums {
		// 验证下标合法性 accNums[i] -k >= 0
		if accNums[i] >= k {
			// 4. 累加每个right 对应(accNums[right] -  k) 的数量
			ans += count[accNums[i]-k]
		}
	}
	return ans
}
