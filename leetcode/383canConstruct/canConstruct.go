package main

import "fmt"

// canConstruct判断ransomNote能否由magazine的数据组成
func canConstruct(ransomNote string, magazine string) bool {
	/*
		这是一道简单题，那么我尽量抛弃可能复杂的解法
		我的解题思路是：直接暴力破解，把magazine转化为map,记录有哪些字符以及字符出现的个数
		然后遍历ransomNote，读到某个字符减一次，直到错误
	*/
	magazineMap := make(map[string]int)
	for _, r := range magazine {
		if _, ok := magazineMap[string(r)]; ok {
			magazineMap[string(r)]++
		} else {
			magazineMap[string(r)] = 1
		}
	}
	for _, r := range ransomNote {
		if _, ok := magazineMap[string(r)]; ok {
			if magazineMap[string(r)] <= 0 {
				return false
			}
			magazineMap[string(r)]--
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("%t\n", canConstruct("aa", "aab"))
}
