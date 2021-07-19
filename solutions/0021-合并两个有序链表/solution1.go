package _1_合并两个有序链表

import . "algorithm-go/utils"

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
