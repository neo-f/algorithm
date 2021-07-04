package week31

// https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/submissions/
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var ans []int
	var pre func(*Node)
	pre = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, c := range node.Children {
			pre(c)
		}
	}
	pre(root)
	return ans
}
