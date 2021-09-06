package week11

import "strconv"

type StringStack []string

func (s *StringStack) Push(r string) {
	*s = append(*s, r)
}

func (s *StringStack) Pop() string {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func calc(b, a, op string) int {
	ib, _ := strconv.Atoi(b)
	ia, _ := strconv.Atoi(a)
	switch op {
	case "+":
		return ia + ib
	case "-":
		return ia - ib
	case "*":
		return ia * ib
	case "/":
		return ia / ib
	}
	return 0
}

func evalRPN(tokens []string) int {
	stack := StringStack{}
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			stack.Push(strconv.Itoa(calc(stack.Pop(), stack.Pop(), token)))
		default:
			stack.Push(token)
		}
	}
	result, _ := strconv.Atoi(stack.Pop())
	return result
}

// 224. 基本计算器    https://leetcode-cn.com/problems/basic-calculator/
// 227. 基本计算器 II https://leetcode-cn.com/problems/basic-calculator-ii/
type RuneStack []rune

func (s *RuneStack) Push(r rune) {
	*s = append(*s, r)
}

func (s *RuneStack) Pop() rune {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func calculate(s string) int {
	// 转换为逆波兰表达式
	var tokens []string
	// 准备压入符号的栈
	ops := RuneStack{}

	// 处理的数值
	var parsingVal int
	// 是否正在处理数字
	var parsingNum bool
	for _, char := range s {
		if char >= '0' && char <= '9' {
			parsingNum = true
			// 从左往右处理
			parsingVal = parsingVal*10 + int(char-'0')
			continue
		} else if parsingNum {
			// 数字处理结束, 存入tokens
			parsingNum = false
			tokens = append(tokens, strconv.Itoa(parsingVal))
			parsingVal = 0
		}
		if char == ' ' {
			continue
		}
		// 符号入栈
		if char == '(' {
			ops.Push(char)
			continue
		}
		if char == ')' {
			// 持续出栈，直到左括号，并把栈内的符号入结果
			for ops.Top() != '(' {
				tokens = append(tokens, string(ops.Pop()))
			}
			// 多一个左括号
			ops.Pop()
			continue
		}
		// 剩下的条件只剩下四种运算符
		if char == '+' || char == '-' || char == '*' || char == '/' {
			// 如果运算优先级比较高的话，那么可以直接入结果，否则压入运算符栈内
			for len(ops) != 0 && getRank(ops.Top()) >= getRank(char) {
				tokens = append(tokens, string(ops.Pop()))
			}
			ops.Push(char)
			continue
		}
	}
	// 遍历结束
	// 检查是否还有处理中的数值
	if parsingNum {
		tokens = append(tokens, strconv.Itoa(parsingVal))
	}
	// 检查栈中剩余的元素
	for len(ops) != 0 {
		tokens = append(tokens, string(ops.Pop()))
	}
	return evalRPN(tokens)
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
