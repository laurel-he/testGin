package main

import "fmt"

// isSubsequence 判断s是否是t的子序列
func isSubsequence(s string, t string) bool {
	/*
		我的解题思路是搞两个指针，分别指向s和t的元素，如果可以把s遍历完就说明是子序列
	*/
	ps := 0
	pt := 0
	if len(s) > len(t) {
		return false
	}
	if len(s) == 0 {
		return true
	}
	for {
		if pt == len(t)-1 {
			// 这里避免一种情况，虽然相等，但是子序列还没有遍历完
			if s[ps] == t[pt] && ps == len(s)-1 {
				return true
			} else {
				return false
			}
		}
		if s[ps] == t[pt] {
			// 可能这已经是子序列最后一个字符了
			if ps == len(s)-1 {
				return true
			}
			ps++
			pt++
		} else {
			pt++
		}
	}
}

func main() {
	fmt.Println(isSubsequence("acb", "ahbgdc"))
}
