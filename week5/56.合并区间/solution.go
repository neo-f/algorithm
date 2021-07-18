package main

import "sort"

type Events [][]int

// Len is the number of elements in the collection.
func (e Events) Len() int {
	return len(e)
}

func (e Events) Less(i int, j int) bool {
	if e[i][0] == e[j][0] {
		return e[i][1] > e[j][1]
	}
	return e[i][0] < e[j][0]
}

func (e Events) Swap(i int, j int) {
	e[i], e[j] = e[j], e[i]
}

func merge(intervals [][]int) [][]int {
	// 差分事件数组
	events := make([][]int, len(intervals)*2)

	for i, interval := range intervals {
		events[i*2] = []int{interval[0], +1}
		events[i*2+1] = []int{interval[1], -1}
	}
	sort.Sort(Events(events))

	var ans [][]int
	count := 0
	recordLeft := 0
	for _, event := range events {
		idx, evt := event[0], event[1]
		// 加之前count=0, 说明开始了新的区间
		if count == 0 {
			recordLeft = idx
		}
		// 累加差异
		count += evt
		// 加完后结果为0，说明事件影响结束，记录答案
		if count == 0 {
			ans = append(ans, []int{recordLeft, idx})
		}
	}
	return ans
}
