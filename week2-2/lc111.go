package week21

import (
	. "algorithm-go/utils"
	"math"
)

// 1. 笨办法
//func minDepth(root *TreeNode) int {
//	ans := math.MaxInt64
//	var traverse func(node *TreeNode, depth int)
//	traverse = func(node *TreeNode, depth int) {
//		// 当前节点为空，说明遍历到头了，深度为上一层的层数，更新答案
//		if node == nil {
//			ans = MinInt(ans, depth-1)
//			return
//		}
//		// 当前节点没有叶子节点，深度为该层层数，更新答案
//		if node.Left == nil && node.Right == nil {
//			ans = MinInt(ans, depth)
//			return
//		}
//		// 左节点为空，右节点不为空，则继续遍历右节点
//		if node.Left == nil {
//			traverse(node.Right, depth+1)
//			return
//		}
//		// 同上
//		if node.Right == nil {
//			traverse(node.Left, depth+1)
//			return
//		}
//		// 左右都不为空
//		traverse(node.Left, depth+1)
//		traverse(node.Right, depth+1)
//	}
//	traverse(root, 1)
//	return ans
//}

// 2. 优化
func minDepth(root *TreeNode) int {
	// 根节点为空
	if root == nil {
		return 0
	}
	// 没有任何叶子节点
	if root.Left == nil && root.Right == nil {
		return 1
	}

	depth := math.MaxInt64
	// 分别计算左右各自的节点, 更新答案
	if root.Left != nil {
		depth = MinInt(depth, minDepth(root.Left))
	}
	if root.Right != nil {
		depth = MinInt(depth, minDepth(root.Right))
	}
	// 有任一子节点，返回节点+1
	return depth + 1
}
