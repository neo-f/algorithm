package week1

// Easy

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

// 26. 删除有序数组中的重复项 https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	fast, slow := 0, 0
	length := len(nums)
	for fast < length {
		if fast == 0 || nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

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

// 206. 反转链表 https://leetcode-cn.com/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

// 21. 合并两个有序链表 https://leetcode-cn.com/problems/merge-two-sorted-lists/
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	curr := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = &ListNode{
				Val:  l1.Val,
				Next: nil,
			}
			l1 = l1.Next
		} else {
			curr.Next = &ListNode{
				Val:  l2.Val,
				Next: nil,
			}
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}
	return head.Next
}

// 66. 加一 https://leetcode-cn.com/problems/plus-one/
// 输入：digits = [1,2,3]
// 输出：[1,2,4]

func plusOne(digits []int) []int {
	ptr := len(digits) - 1
	plus := true
	for ptr >= 0 {
		if digits[ptr] == 9 && plus {
			// 如果是9且需要进位, 则数字改为0, 保持进位
			digits[ptr] = 0
		} else if plus {
			// 如果不是9且需要进位, 则数字+1, 不再进位
			digits[ptr]++
			plus = false
		} else {
			// 如果不是9且不需要进位, 不动, 且不再进位
			plus = false
		}
		ptr--
	}
	// 如果结果仍然需要进位，则头补1
	if plus {
		digits = append([]int{1}, digits...)
	}
	return digits
}
