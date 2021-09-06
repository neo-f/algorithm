package week31

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/submissions/
func buildTree(inorder []int, postorder []int) *TreeNode {
	// inorder   左中右，
	// postorder 左右中,
	if len(postorder) == 0 {
		return nil
	}

	// 1. 找到根，根永远是postorder的最后一个元素
	rootVal := postorder[len(postorder)-1]
	// 2. 找到根在中序数组中的位置 (此处可以使用HashMap优化一下, 因为题目中说明值不重复)
	rootIdx := 0
	for inorder[rootIdx] != rootVal {
		rootIdx++
	}
	// 3. 递归同105
	node := &TreeNode{Val: inorder[rootIdx]}
	node.Left = buildTree(inorder[:rootIdx], postorder[:rootIdx])
	node.Right = buildTree(inorder[rootIdx+1:], postorder[rootIdx:len(postorder)-1])
	return node
}
