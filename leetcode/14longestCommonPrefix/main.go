package main

import "fmt"

// longestCommonPrefix获取最长公共前缀
func longestCommonPrefix(strs []string) string {
	/*
		我的解题思路:
		一个指针，指向最长公共前缀的位置，若是公共前缀，指针往下走
	*/
	point := 0
	for {
		if len(strs) == 0 {
			return "" // 注意处理边界条件
		}
		if strs[0] == "" {
			return ""
		}
		if len(strs[0]) <= point {
			return strs[0][:point]
		}
		commonRune := strs[0][point]
		for _, str := range strs {
			if len(str) > point {
				if str[point] != commonRune {
					return str[:point]
				}
			} else {
				return str[:point]
			}
		}
		point++
	}
}

func main() {
	res := longestCommonPrefix([]string{"ab", "a"})
	fmt.Println(res)
}
