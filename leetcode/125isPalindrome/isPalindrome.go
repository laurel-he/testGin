package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	/*
		直接搞两个指针，指向第一个和最后一个，直到指针相遇
	*/
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	i := 0
	j := len(s) - 1
	for {
		if i < j {
			// 注意，题目里写的是移除所有非字母数字字符
			if !unicode.IsLetter(rune(s[i])) && !unicode.IsNumber(rune(s[i])) {
				i++
				continue
			}
			if !unicode.IsLetter(rune(s[j])) && !unicode.IsNumber(rune(s[j])) {
				j--
				continue
			}
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		} else {
			return true
		}
	}
}

func main() {
	//fmt.Println(isPalindrome("abcba"))
	fmt.Println(isPalindrome("0P"))
}
