package main

import "fmt"

// removeDuplicates 删除数组中的重复项
func removeDuplicates(nums []int) int {
	/*
		题目描述中依然提到要注意原地删除，即不引入额外的内存空间
		注意到描述中有提到是按升序排列的，即每一个元素只需要与前一个元素对比，无需全部对比
		我的解题思路是，一个变量记录当前值，每次去对比
		（1）若出现重复就删掉，循环时下标会往前移动
	*/
	if len(nums) <= 1 {
		return len(nums)
	}
	now := nums[0]
	nowPoint := 1 // 当前指向第几个元素，作比较的基础数字从0开始，被比较的从第一个开始
	// 0,1,1,1,2,2,3,3,4
	for {
		if nowPoint > len(nums)-1 {
			// 不需要管删除了多少，反正slice会自动往前移动，只需要知道已经读到最后一个就行了
			break
		}
		if now == nums[nowPoint] {
			nums = append(nums[:nowPoint], nums[nowPoint+1:]...) // 删掉nowPoint
			// 每删掉一个，nums长度-1，下一个元素补位
		} else {
			now = nums[nowPoint]
			nowPoint++
		}
	}
	return len(nums)
}

func main() {
	numLen := removeDuplicates([]int{1, 1})
	fmt.Println(numLen)
}
