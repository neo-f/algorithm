package _8_合并两个有序数组

// 88. 合并两个有序数组 https://leetcode-cn.com/problems/merge-sorted-array/submissions/
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	ptr := m + n - 1
	for ptr >= 0 {
		switch {
		// nums1到头
		case p1 == -1:
			nums1[ptr] = nums2[p2]
			p2--
		// nums2到头
		case p2 == -1:
			nums1[ptr] = nums1[p1]
			p1--
		case nums1[p1] > nums2[p2]:
			nums1[ptr] = nums1[p1]
			p1--
		default:
			nums1[ptr] = nums2[p2]
			p2--
		}
		ptr--
	}
}
