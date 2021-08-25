package main

type NumArray struct {
	// pre_sums
	data []int
}

func Constructor(nums []int) NumArray {
	arr := NumArray{data: make([]int, len(nums)+1)}
	for i := range nums {
		arr.Update(i, nums[i])
	}
	return arr
}

func (arr *NumArray) Update(index int, val int) {
	delta := val - arr.SumRange(index, index)
	index++
	for ; index <= len(arr.data); index += index & -index {
		arr.data[index] += delta
	}
}

func (arr *NumArray) sum(i int) int {
	var ans int
	for ; i > 0; i -= i & -i {
		ans += arr.data[i]
	}
	return ans
}

func (arr *NumArray) SumRange(left int, right int) int {
	//right++
	//left++
	//return arr.sum(right) - arr.sum(left-1)
	return arr.sum(right+1) - arr.sum(left)
}
