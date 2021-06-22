package week11

import . "algorithm-go/utils"

// 141. 环形链表 https://leetcode-cn.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		head = head.Next
		if fast == head {
			return true
		}
	}
	return false
}
