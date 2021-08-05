package _684_冗余连接

type disjointSet struct {
	data []int
}

func newDSet(n int) *disjointSet {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}
	return &disjointSet{data: data}
}

func (s *disjointSet) Find(x int) int {
	if s.data[x] == x {
		return x
	}
	s.data[x] = s.Find(s.data[x])
	return s.data[x]
}
func (s *disjointSet) Union(x, y int) {
	x = s.Find(x)
	y = s.Find(y)
	if x != y {
		s.data[x] = y
	}
}

func findRedundantConnection(edges [][]int) []int {
	// 有边+1个顶点
	disjointSet := newDSet(len(edges)+1)
	for _, edge := range edges{
		x, y := disjointSet.Find(edge[0]), disjointSet.Find(edge[1])
		if x != y{
			disjointSet.Union(x, y)
		}else{
			return edge
		}
	}
	return nil
}
