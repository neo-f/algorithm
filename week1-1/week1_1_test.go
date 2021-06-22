package week11

import (
	"algorithm-go/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 3, 5, 0, 0, 0}
	nums2 := []int{2, 4, 6}
	merge(nums1, 3, nums2, 3)
	assert.Equal(t, nums1, []int{1, 2, 3, 4, 5, 6})

	nums3 := []int{2, 2, 2, 4, 0, 0, 0}
	nums4 := []int{2, 4, 6}
	merge(nums3, 4, nums4, 3)
	assert.Equal(t, nums3, []int{2, 2, 2, 2, 4, 4, 6})

	nums5 := make([]int, 0)
	nums6 := make([]int, 0)
	merge(nums5, 0, nums6, 0)
	assert.Equal(t, nums5, []int{})
}

func TestRemoveDuplicates(t *testing.T) {
	case1 := []int{1, 1, 2, 2, 3}
	result1 := removeDuplicates(case1)
	assert.Equal(t, result1, 3)
	assert.Equal(t, case1[:3], []int{1, 2, 3})

	case2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result2 := removeDuplicates(case2)
	assert.Equal(t, result2, 5)
	assert.Equal(t, case2[:5], []int{0, 1, 2, 3, 4})
}

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	assert.Equal(t, nums, []int{1, 3, 12, 0, 0})
}

func TestReverseList(t *testing.T) {
	array := utils.SliceToList([]int{1, 2, 3})
	assert.Equal(t, reverseList(array), utils.SliceToList([]int{3, 2, 1}))
}

func TestMergeTwoLists(t *testing.T) {
	l1 := utils.SliceToList([]int{1, 3, 5})
	l2 := utils.SliceToList([]int{2, 4, 6})
	r := mergeTwoLists(l1, l2)
	t.Log(l1, l2, r)
	assert.Equal(t, r, utils.SliceToList([]int{1, 2, 3, 4, 5, 6}))
}

func TestPlusOne(t *testing.T) {
	assert.Equal(t, plusOne([]int{1, 2, 3}), []int{1, 2, 4})
	assert.Equal(t, plusOne([]int{1, 2, 9}), []int{1, 3, 0})
	assert.Equal(t, plusOne([]int{9}), []int{1, 0})
}

func TestIsValid(t *testing.T) {
	assert.Equal(t, isValid("{}"), true)
	assert.Equal(t, isValid("{[]}"), true)
	assert.Equal(t, isValid("{[]})"), false)
	assert.Equal(t, isValid("{[{]}"), false)
}

func TestMinStack(t *testing.T) {
	s := NewMinStack()
	s.Push(1)
	s.Push(2)
	s.Push(-10)
	s.Push(2)
	assert.Equal(t, s.Top(), 2)
	assert.Equal(t, s.GetMin(), -10)
	s.Pop()
	assert.Equal(t, s.Top(), -10)
	assert.Equal(t, s.GetMin(), -10)
	s.Pop()
	assert.Equal(t, s.Top(), 2)
	assert.Equal(t, s.GetMin(), 1)
}

func TestEvalRPN(t *testing.T) {
	assert.Equal(t, evalRPN([]string{"2", "1", "+", "3", "*"}), 9)
	assert.Equal(t, evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}), 22)
}
func TestCircularDeque(t *testing.T) {
	dq := NewMyCircularDeque(3)
	assert.Equal(t, dq.InsertLast(1), true)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
	assert.Equal(t, dq.InsertLast(2), true)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
	assert.Equal(t, dq.InsertFront(3), true)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
	assert.Equal(t, dq.InsertFront(4), false)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
	assert.Equal(t, dq.GetRear(), 2)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
	assert.Equal(t, dq.IsFull(), true)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
	assert.Equal(t, dq.DeleteLast(), true)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
	assert.Equal(t, dq.InsertFront(4), true)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
	assert.Equal(t, dq.GetFront(), 4)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
}

func TestCalculate(t *testing.T) {
	assert.Equal(t, calculate("1+2*3"), 7)
	assert.Equal(t, calculate("(1+2)*3"), 9)
}

func TestSubarraySum(t *testing.T) {
	assert.Equal(t, subarraySum([]int{1, 1, 1}, 2), 2)
}

func TestReverseKGroup(t *testing.T) {
	c1 := utils.SliceToList([]int{1, 2, 3, 4, 5, 6})
	e1 := utils.SliceToList([]int{2, 1, 4, 3, 6, 5})
	assert.Equal(t, reverseKGroup(c1, 2), e1)
	c2 := utils.SliceToList([]int{1, 2, 3, 4, 5, 6})
	e2 := utils.SliceToList([]int{4, 3, 2, 1, 5, 6})
	assert.Equal(t, reverseKGroup(c2, 4), e2)
	c3 := utils.SliceToList([]int{1, 2, 3, 4, 5})
	e3 := utils.SliceToList([]int{2, 1, 4, 3, 5})
	assert.Equal(t, reverseKGroup(c3, 2), e3)
	c4 := utils.SliceToList([]int{1, 2, 3, 4, 5})
	e4 := utils.SliceToList([]int{1, 2, 3, 4, 5})
	assert.Equal(t, reverseKGroup(c4, 1), e4)
}
