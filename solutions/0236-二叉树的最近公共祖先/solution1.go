package week31

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parentMap := make(map[*TreeNode]*TreeNode, 0)

	var traverse func(*TreeNode, *TreeNode)
	// 遍历一遍二叉树，并记录各节点的父节点
	traverse = func(parent, node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parentMap[node.Left] = node
			traverse(node, node.Left)
		}
		if node.Right != nil {
			parentMap[node.Right] = node
			traverse(node, node.Right)
		}
	}
	traverse(nil, root)

	// 标记p的父级路径
	pathSet := make(map[*TreeNode]struct{}, 0)
	for p != nil {
		pathSet[p] = struct{}{}
		p = parentMap[p]
	}
	for q != nil {
		if _, ok := pathSet[q]; ok {
			return q
		}
		q = parentMap[q]
	}
	return nil
}
