package main

import "strings"

func lengthOfLastWord(s string) int {
	/**
	这道题直接去除字符串首尾的空格，然后用空格拆分，返回最后一个即可
	*/
	s = strings.TrimSpace(s)
	allStr := strings.Split(s, " ")
	return len(allStr[len(allStr)-1])
}

func main() {

}
