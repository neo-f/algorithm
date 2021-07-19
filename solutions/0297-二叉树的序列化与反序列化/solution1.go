package week31

import (
	. "algorithm-go/utils"
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/

type Codec struct {
}

func NewCodec() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	var tokens []string

	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			tokens = append(tokens, "nil")
			return
		}
		tokens = append(tokens, strconv.Itoa(node.Val))
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)

	return strings.Join(tokens, ",")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	tokens := strings.Split(data, ",")
	var idx = 0
	var parse func() *TreeNode

	parse = func() *TreeNode {
		token := tokens[idx]
		if token == "nil" {
			idx++
			return nil
		}
		val, _ := strconv.Atoi(token)
		idx++
		node := &TreeNode{Val: val}
		node.Left = parse()
		node.Right = parse()
		return node
	}
	return parse()
}
