package week1

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var ss []string
	nl := l
	for nl != nil {
		ss = append(ss, strconv.Itoa(nl.Val))
		nl = nl.Next
	}
	return strings.Join(ss, "->")
}

// SliceToList 数组切片转单链表
func SliceToList(slice []int) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	curr := head
	for i := range slice {
		curr.Next = &ListNode{
			Val:  slice[i],
			Next: nil,
		}
		curr = curr.Next
	}
	return head.Next
}
