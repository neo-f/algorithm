package week1

import "strconv"

// 25. K 个一组翻转链表 https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	// TODO
	return nil
}

// 224. 基本计算器    https://leetcode-cn.com/problems/basic-calculator/
// 227. 基本计算器 II https://leetcode-cn.com/problems/basic-calculator-ii/
func calculate(s string) int {
	// TODO: 多复习
	var rpnTokens []string
	opts := runeStack{}
	parsingNum := false
	v := 0
	for _, token := range s {
		if token >= '0' && token <= '9' {
			parsingNum = true
			v = v*10 + int(token-'0')
			continue
		} else if parsingNum {
			// 处理完数值，并且已经结束
			parsingNum = false
			rpnTokens = append(rpnTokens, strconv.Itoa(v))
			v = 0
		}
		if token == ' ' {
			continue
		}
		if token == '(' {
			opts.Push(token)
			continue
		}
		if token == ')' {
			for opts.Top() != '(' {
				rpnTokens = append(rpnTokens, string(opts.Top()))
				opts.Pop()
			}
			opts.Pop()
			continue
		}
		for len(opts) != 0 && getRank(opts.Top()) >= getRank(token) {
			rpnTokens = append(rpnTokens, string(opts.Top()))
			opts.Pop()
		}
		opts.Push(token)
	}
	if parsingNum {
		rpnTokens = append(rpnTokens, strconv.Itoa(v))
	}
	for len(opts) != 0 {
		o, _ := opts.Pop()
		rpnTokens = append(rpnTokens, string(o))
	}
	return evalRPN(rpnTokens)
}

func getRank(token rune) int {
	switch token {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}
