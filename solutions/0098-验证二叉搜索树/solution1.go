package week21

import (
	. "algorithm-go/utils"
	"math"
)

func isValidBST(root *TreeNode) bool {
	var isValid func(node *TreeNode, min, max int) bool
	// min 保证合法的最小值
	// max 保证合法的最大值
	isValid = func(node *TreeNode, min, max int) bool {
		// 终止条件：空节点一定合法
		if node == nil {
			return true
		}
		// 不合法条件: 1. 值比最小值还小 2. 值比最大值还大
		if node.Val <= min || node.Val >= max {
			return false
		}
		// 递归
		// 左节点：最小值仍为最小值，最大值不超过父节点的值
		// 右节点：最大值仍为最大值，最小值不小于父节点的值
		return isValid(node.Left, min, node.Val) && isValid(node.Right, node.Val, max)
	}
	return isValid(root, math.MinInt64, math.MaxInt64)
}
