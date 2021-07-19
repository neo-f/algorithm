package week31

import . "algorithm-go/utils"

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
func inorderTraversal(root *TreeNode) []int {
	var ans []int

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(root.Left)
		ans = append(ans, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return ans
}
