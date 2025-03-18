package main

import (
	"fmt"
	"sort"
)

// merge 合并nums1和nums2这两个有序数组，m和n分别表示数组长度，注意nums1的长度实际是m+n，多余的填充0
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 我的解题思路是直接强行合并，然后冒泡排序
	allSlice := append(nums1[:m], nums2...)
	sort.Ints(allSlice)
	fmt.Println(allSlice)
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
