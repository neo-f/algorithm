package main

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type SegmentTree struct {
	tree []int
	lazy []int
	N, H int
}

func NewSegmentTree(n int) *SegmentTree {
	h := 1
	for 1<<h < n {
		h++
	}
	return &SegmentTree{
		N:    n,
		H:    h,
		tree: make([]int, n*4),
		lazy: make([]int, n),
	}
}

func (t *SegmentTree) apply(i int, val int) {
	t.tree[i] = max(t.tree[i], val)
	if i < t.N {
		t.lazy[i] = max(t.lazy[i], val)
	}
}

func (t *SegmentTree) pull(i int) {
	for i > 1 {
		i >>= 1
		t.tree[i] = max(t.tree[i*2], t.tree[i*2+1])
		t.tree[i] = max(t.tree[i], t.lazy[i])
	}
}

func (t *SegmentTree) push(i int) {
	for h := t.H; h > 0; h-- {
		if y := i >> h; t.lazy[y] > 0 {
			t.apply(y*2, t.lazy[y])
			t.apply(y*2+1, t.lazy[y])
			t.lazy[y] = 0
		}
	}
}
func (t *SegmentTree) Update(l, r, h int) {
	l += t.N
	r += t.N
	l0, r0 := l, r
	for l <= r {
		if l&1 == 1 {
			t.apply(l, h)
			l++
		}
		if r&1 == 0 {
			t.apply(r, h)
			r--
		}
		l >>= 1
		r >>= 1
	}
	t.pull(l0)
	t.pull(r0)
}

func (t *SegmentTree) Query(l, r int) int {
	l += t.N
	r += t.N
	t.push(l)
	t.push(r)
	var ans int
	for l <= r {
		if l&1 == 1 {
			ans = max(ans, t.tree[l])
			l++
		}
		if r&1 == 0 {
			ans = max(ans, t.tree[r])
			r--
		}
		l >>= 1
		r >>= 1
	}
	return ans
}

func fallingSquares(positions [][]int) []int {
	// set
	coord := make(map[int]struct{})
	for _, pos := range positions {
		coord[pos[0]] = struct{}{}
		coord[pos[0]+pos[1]-1] = struct{}{}
	}
	// set -> list
	sortedCoord := make([]int, 0, len(coord))
	for k := range coord {
		sortedCoord = append(sortedCoord, k)
	}
	sort.Ints(sortedCoord)

	index := make(map[int]int)
	for idx, coord := range sortedCoord {
		index[coord] = idx
	}

	tree := NewSegmentTree(len(sortedCoord))
	var best int
	var ans []int
	for _, pos := range positions {
		l := index[pos[0]]
		r := index[pos[0]+pos[1]-1]
		h := tree.Query(l, r) + pos[1]
		tree.Update(l, r, h)
		best = max(best, h)
		ans = append(ans, best)
	}
	return ans
}
