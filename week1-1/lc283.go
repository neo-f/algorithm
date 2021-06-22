package week11

// 283. 移动零 https://leetcode-cn.com/problems/move-zeroes/
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	length := len(nums)
	for fast < length {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	for slow < length {
		nums[slow] = 0
		slow++
	}
}