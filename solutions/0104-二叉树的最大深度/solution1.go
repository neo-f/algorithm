package week21

import (
	"math"
)

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	ans := math.MinInt64

	var find func(node *TreeNode, depth int)
	find = func(node *TreeNode, depth int) {
		// 终止条件：节点为空，层数为上一层
		if node == nil {
			ans = MaxInt(ans, depth-1)
			return
		}
		find(node.Left, depth+1)
		find(node.Right, depth+1)
	}
	find(root, 1)
	return ans
}
