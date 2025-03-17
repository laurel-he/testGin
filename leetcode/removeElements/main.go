package main

import "fmt"

// removeElement 移除元素
func removeElement(nums []int, val int) int {
	/*
		在题目要求中加黑提醒：需要原地移除元素，即不能引入额外的内存空间
		移除元素后，所有元素都要往前移动
		双指针：
		一个指向第一个元素一个指向最后一个元素
		当匹配到数据时，交换两个指针指向数据，并将左指针向右，右指针向左移动
		记录交换的次数，移除
	*/
	p1, p2 := 0, len(nums)
	removeNums := 0
	for {
		if p1 >= p2 {
			nums = nums[:len(nums)-removeNums]
			fmt.Println("nums:", nums)
			//return removeNums
			// 这里有点混乱，题目描述中说返回删除的数量，但实际后面又说返回k为nums中不等于val的数量
			return len(nums)
		}
		if nums[p1] == val {
			// 如果右指针指向的数据也等于val,右指针向左移动
			for {
				// 处理越界情况
				if p1 >= p2 {
					nums = nums[:len(nums)-removeNums]
					return len(nums)
				}
				if nums[p2-1] == val {
					removeNums++
					p2--
				} else {
					break
				}
			}
			nums[p1], nums[p2-1] = nums[p2-1], nums[p1] // 交换
			removeNums++
			p1++
			p2--
		} else {
			// 如果不相等，左指针向右移动
			p1++
		}
	}
}

func main() {
	nums := []int{1}
	val := 1
	num := removeElement(nums, val)
	fmt.Println(num)
}
