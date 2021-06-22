package week11

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

// rune栈
type runeStack []rune

func (s *runeStack) Push(r rune) {
	*s = append(*s, r)
}
func (s *runeStack) Top() rune {
	if len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}
func (s *runeStack) Pop() (r rune, ok bool) {
	length := len(*s)
	if length == 0 {
		return 0, false
	}
	r = (*s)[length-1]
	*s = (*s)[:length-1]
	return r, true
}

type stringStack []string

func (s *stringStack) Push(r string) {
	*s = append(*s, r)
}
func (s *stringStack) Pop() string {
	idx := len(*s) - 1
	r := (*s)[idx]
	*s = (*s)[:idx]
	return r
}
