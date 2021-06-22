package week11

import . "algorithm-go/utils"

// 20. 有效的括号 https://leetcode-cn.com/problems/valid-parentheses/
func isValid(s string) bool {
	stack := RuneStack{}
	for _, token := range s {
		switch token {
		case '(', '[', '{':
			stack.Push(token)
		case ')':
			if stack.Pop() != '(' {
				return false
			}
		case ']':
			if stack.Pop() != '[' {
				return false
			}
		case '}':
			if stack.Pop() != '{' {
				return false
			}
		}
	}

	return len(stack) == 0
}
