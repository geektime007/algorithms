package main

import "fmt"

func isValid1(s string) bool {
	// 空字符串直接返回 true
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '[') || (v == '(') || (v == '{') {
			stack = append(stack, v)
		} else if ((v == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			((v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((v == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println("vim-go")
	var s string

	s = "()"
	fmt.Println(s, " => ", isValid1(s))

	s = "()[]{}"
	fmt.Println(s, " => ", isValid1(s))

	s = "(]"
	fmt.Println(s, " => ", isValid1(s))

	s = "({)]"
	fmt.Println(s, " => ", isValid1(s))

	s = "{[]}"
	fmt.Println(s, " => ", isValid1(s))
}
