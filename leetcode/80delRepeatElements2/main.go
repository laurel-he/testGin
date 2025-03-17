package main

import "fmt"

// removeDuplicates 删除数组中的重复项，使得最多重复两次
func removeDuplicates(nums []int) int {
	/*
		我刚做完编号为26的删除数组中重复项的题目，顺着上一道题的思路：
		依然是有序的，即重复的数据是挨着的，那么在上一道题目的基础上做出简单的修改即可：
		首先：即使重复，指针依旧往下走，直到第二次重复，才认为需要删除
		需要一个额外的变量记录是否已经重复过
	*/
	if len(nums) <= 1 {
		return len(nums)
	}
	now := nums[0]
	nowPoint := 1 // 当前指向第几个元素，作比较的基础数字从0开始，被比较的从第一个开始
	isRepeat := false
	// 0,1,1,1,2,2,3,3,4
	for {
		if nowPoint > len(nums)-1 {
			// 不需要管删除了多少，反正slice会自动往前移动，只需要知道已经读到最后一个就行了
			break
		}
		if now == nums[nowPoint] {
			if isRepeat {
				nums = append(nums[:nowPoint], nums[nowPoint+1:]...) // 删掉nowPoint
			} else {
				isRepeat = true
				nowPoint++
			}
			// 每删掉一个，nums长度-1，下一个元素补位
		} else {
			now = nums[nowPoint]
			nowPoint++
			isRepeat = false
		}
	}
	return len(nums)
}

func main() {
	numLen := removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3})
	fmt.Println(numLen)
}
