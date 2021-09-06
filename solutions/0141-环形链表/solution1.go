package week11

type ListNode struct {
	Next *ListNode
	Val  int
}

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
