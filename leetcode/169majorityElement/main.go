package main

import "fmt"

// majorityElement 数量为n，出现次数大于n/2的元素组成的slice
func majorityElement(nums []int) int {
	/*
		最简单的方法，直接记录每个元素出现的次数，记录到map中，每出现一次+1，只要发现大于n/2立马返回
	*/
	maxNum := len(nums) / 2
	repeatTimeMap := make(map[int]int) // 拿到每个元素出现的次数
	for i := 0; i < len(nums); i++ {
		if _, exist := repeatTimeMap[nums[i]]; !exist {
			repeatTimeMap[nums[i]] = 1
		} else {
			repeatTimeMap[nums[i]]++
		}
		if repeatTimeMap[nums[i]] > maxNum {
			return nums[i]
		}
	}
	return 0
}

func main() {
	res := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	fmt.Println(res)
}
