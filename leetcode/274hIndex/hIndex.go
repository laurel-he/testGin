package main

import (
	"fmt"
	"sort"
)

// hIndex 计算"h指数"
func hIndex(citations []int) int {
	/**
	题目不是很好理解，我理解为：len(citations)=n,存在一个最大的h，使得至少有h个>=h的数字出现，h<=n
	例如：[3,0,6,1,5]，长度为5，存在3，使得至少有3个大于3的数字出现
	我直接先排序，按照从大到小来排，然后找最大的是否是成立的"h"，不成立就找下一个，直到找到成立的位置
	*/
	var h int
	sort.Ints(citations)
	for i := len(citations) - 1; i >= 0; i-- {
		if checkIsH(citations, citations[i]) {
			h = citations[i]
		}
	}
	// 以上循环检测过后，有额外情况，数据不在这个数组范围内，例如[100]的h为1，那么使用长度来判断
	for i := len(citations); i >= 0; i-- {
		if checkIsH(citations, i) {
			h = max(h, i)
		}
	}
	return h
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// checkIsH 检验一个数是否是h
func checkIsH(citations []int, h int) bool {
	canhNum := 0
	for i := 0; i < len(citations); i++ {
		if citations[i] >= h {
			canhNum++
		}
	}
	if canhNum >= h {
		return true
	}
	return false
}

func main() {
	fmt.Println(hIndex([]int{0, 0, 2}))
}
