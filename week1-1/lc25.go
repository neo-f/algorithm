package week11

// 25. K 个一组翻转链表 https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	protectNode := &ListNode{Next: head}
	for currGroupHead, prevGroupHead := head, protectNode; currGroupHead != nil; {
		currGroupEnd := getEnd(currGroupHead, k)
		if currGroupEnd == nil {
			break
		}
		nextGroupHead := currGroupEnd.Next
		// 此处的head和end已经调换，原来的end 变成了新的 head, 所以原来的end.Next会丢失，所以在上一步要记录一
		reverseGroupInner(currGroupHead, currGroupEnd)
		// 上一组的Next指向新的head
		prevGroupHead.Next = currGroupEnd
		// 新的end的Next指向下一组
		currGroupHead.Next = nextGroupHead
		// 记录上一组
		prevGroupHead = currGroupHead
		// 移动当前head到下一组
		currGroupHead = nextGroupHead
	}
	return protectNode.Next
}

func reverseGroupInner(head *ListNode, end *ListNode) {
	originNext := end.Next
	var prev *ListNode
	for head != originNext {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
}

func getEnd(head *ListNode, k int) *ListNode {
	h := head
	for h != nil {
		k--
		if k == 0 {
			return h
		}
		h = h.Next
	}
	return h
}
