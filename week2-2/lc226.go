package week21

import . "algorithm-go/utils"

func invertTree(root *TreeNode) *TreeNode {
	var invert func(root *TreeNode)
	invert = func(node *TreeNode) {
		if node == nil || (node.Left == nil && node.Right == nil) {
			return
		}
		invert(node.Left)
		invert(node.Right)
		node.Left, node.Right = node.Right, node.Left
	}
	invert(root)
	return root
}
