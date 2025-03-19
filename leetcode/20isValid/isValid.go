package main

import "fmt"

// isValid 判断括号是否有效
func isValid(s string) bool {
	/*
		判断括号是否有效：用栈判断，当是'('时入栈，遇到')'出栈，如果不出错那就成立
	*/
	checkMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	// 不初始化长度，需要用到这个长度来判断
	var checkStack []rune
	for _, strItem := range s {
		switch strItem {
		case '(', '[', '{':
			checkStack = append(checkStack, strItem)
		case ')', ']', '}':
			if len(checkStack) <= 0 {
				return false
			}
			if checkMap[strItem] != checkStack[len(checkStack)-1] {
				return false
			} else {
				checkStack = checkStack[:len(checkStack)-1]
			}
		}
	}
	// 额外判断，栈不为空，那就不成立
	if len(checkStack) > 0 {
		fmt.Println("checkStack", checkStack)
		return false
	}
	return true
}

func main() {
	fmt.Println("isValid", isValid("()[]{}"))
}
