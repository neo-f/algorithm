package week11

// 20. 有效的括号 https://leetcode-cn.com/problems/valid-parentheses/
func isValid(s string) bool {
	stack := runeStack{}
	for _, token := range s {
		switch token {
		case '(', '[', '{':
			stack.Push(token)
		case ')':
			if r, ok := stack.Pop(); !ok || r != '(' {
				return false
			}
		case ']':
			if r, ok := stack.Pop(); !ok || r != '[' {
				return false
			}
		case '}':
			if r, ok := stack.Pop(); !ok || r != '{' {
				return false
			}
		}
	}

	return len(stack) == 0
}
