package week31

// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/

type NodeInfo struct {
	Depth int
	Node  *Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	var queue []*NodeInfo

	queue = append(queue, &NodeInfo{Depth: 0, Node: root})

	for len(queue) != 0 {
		node := queue[0]
		// Pop 队头
		queue = queue[1:]
		// 初始化检查ans的大小，如果空间不够可能造成索引越界，则扩展一次ans
		// 因为Depth每次都只加1，所以只需要扩展一个
		if len(ans) <= node.Depth {
			ans = append(ans, make([]int, 0))
		}
		ans[node.Depth] = append(ans[node.Depth], node.Node.Val)
		for _, c := range node.Node.Children {
			queue = append(queue, &NodeInfo{
				Depth: node.Depth + 1,
				Node:  c,
			})
		}
	}

	return ans
}
