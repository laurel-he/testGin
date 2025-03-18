package main

import (
	"fmt"
	"strings"
)

// wordPattern 判断是否遵从某种规律
func wordPattern(pattern string, s string) bool {
	/*
		这道题和我刚做过的205非常相似，只不过从单个的单词变成了词汇，因此我使用相同的解法
		如果允许a-<dog且b->dog，这道题甚至不需要反向再判断，我一开始也是这么以为的
		但实际上不允许，因此还是需要反向判断
	*/
	checkMap := make(map[byte]string, len(pattern))
	checkMapReturn := make(map[string]byte, len(pattern))
	sSlice := strings.Split(s, " ")
	if len(sSlice) != len(pattern) {
		// 长度都不相等，直接不匹配
		return false
	}
	for i := 0; i < len(sSlice); i++ {
		if _, ok := checkMap[pattern[i]]; !ok {
			checkMap[pattern[i]] = sSlice[i]
		} else {
			if sSlice[i] != checkMap[pattern[i]] {
				return false
			}
		}
		if _, ok := checkMapReturn[sSlice[i]]; !ok {
			checkMapReturn[sSlice[i]] = pattern[i]
		} else {
			if pattern[i] != checkMapReturn[sSlice[i]] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}
