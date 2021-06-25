package utils

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

// SliceToList 数组切片转单链表.
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

// RuneStack rune栈
type RuneStack []rune

func (s *RuneStack) Push(r rune) {
	*s = append(*s, r)
}
func (s *RuneStack) Top() rune {
	if len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}
func (s *RuneStack) Pop() rune {
	length := len(*s)
	if length == 0 {
		return 0
	}
	r := (*s)[length-1]
	*s = (*s)[:length-1]
	return r
}

// StringStack string栈.
type StringStack []string

func (s *StringStack) Push(r string) {
	*s = append(*s, r)
}
func (s *StringStack) Top() string {
	if len(*s) == 0 {
		return ""
	}
	return (*s)[len(*s)-1]
}
func (s *StringStack) Pop() string {
	length := len(*s)
	if length == 0 {
		return ""
	}
	r := (*s)[length-1]
	*s = (*s)[:length-1]
	return r
}

type IntStack []int

func (s *IntStack) Push(r int) {
	*s = append(*s, r)
}
func (s *IntStack) Top() int {
	if len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}
func (s *IntStack) IsEmpty() bool {
	return len(*s) == 0
}
func (s *IntStack) Pop() int {
	length := len(*s)
	if length == 0 {
		return 0
	}
	r := (*s)[length-1]
	*s = (*s)[:length-1]
	return r
}

func MaxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func MinInt(i, j int) int {
	if i < j {
		return i
	}
	return j
}
