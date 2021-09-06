package week31

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	node := &TreeNode{Val: preorder[0]}

	// 找到的根在前序数组的下标即左段的长度
	leftSize := 0
	for preorder[0] != inorder[leftSize] {
		leftSize++
	}

	node.Left = buildTree1(preorder[1:1+leftSize], inorder[:leftSize])
	node.Right = buildTree1(preorder[leftSize+1:], inorder[leftSize+1:])
	return node
}
