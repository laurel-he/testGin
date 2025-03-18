package main

import "fmt"

// isIsomorphic 是否t中的字符串是按照s中一样的顺序出现
func isIsomorphic(s string, t string) bool {
	/*
		这道题关键是出现的“位置”以及“次数”是相等的
		例如"egg"和"add"就是“同构字符串”
		可以定义一个map，其中列入map[e]=a,看是否能完全成立
		需要额外注意，还要反过来算一遍，因为可能出现：
		s="badc" t="baba"这种情况，这是map[d]=b会额外创建map，但反之不成立
	*/
	omorphiMapS := make(map[byte]byte)
	omorphiMapT := make(map[byte]byte)
	if len(s) != len(t) {
		// 字符串长度都不相等的情况首先排除
		return false
	}
	for i := 0; i < len(s); i++ {
		if _, ok := omorphiMapS[s[i]]; !ok {
			omorphiMapS[s[i]] = t[i]
		} else {
			if omorphiMapS[s[i]] != t[i] {
				return false
			}
		}
		if _, ok := omorphiMapT[t[i]]; !ok {
			omorphiMapT[t[i]] = s[i]
		} else {
			if omorphiMapT[t[i]] != s[i] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("badc", "baba"))
}
