package week1

import (
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
	array := SliceToList([]int{1, 2, 3})
	assert.Equal(t, reverseList(array), SliceToList([]int{3, 2, 1}))
}

func TestMergeTwoLists(t *testing.T) {
	l1 := SliceToList([]int{1, 3, 5})
	l2 := SliceToList([]int{2, 4, 6})
	r := mergeTwoLists(l1, l2)
	t.Log(l1, l2, r)
	assert.Equal(t, r, SliceToList([]int{1, 2, 3, 4, 5, 6}))
}

func TestPlusOne(t *testing.T) {
	assert.Equal(t, plusOne([]int{1, 2, 3}), []int{1, 2, 4})
	assert.Equal(t, plusOne([]int{1, 2, 9}), []int{1, 3, 0})
	assert.Equal(t, plusOne([]int{9}), []int{1, 0})
}
