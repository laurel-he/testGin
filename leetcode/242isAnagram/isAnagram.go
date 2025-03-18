package main

import "fmt"

// isAnagram 是否是有效的“字母异位词”
func isAnagram(s string, t string) bool {
	/*
		所谓的“字母异位词”是指所有词汇重新排列，且只使用一次，就能构成另一个词汇
		我不久前刚做了一道类似的题目，map记录词汇出现的次数，判断另一个是否满足
	*/
	checkMap := make(map[rune]int)
	if len(s) != len(t) {
		return false
	}
	for _, v := range s {
		if _, ok := checkMap[v]; ok {
			checkMap[v]++
		} else {
			checkMap[v] = 1
		}
	}
	for _, v := range t {
		if _, ok := checkMap[v]; ok {
			checkMap[v]--
			if checkMap[v] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("rat", "car"))
}
