package main

import "fmt"

// jump 跳跃游戏：获得最少跳跃次数
func jump(nums []int) int {
	/*
		我自己尝试解决，10分钟后没有任何思路……
		看了官方解法，我理解是从最后一个元素倒推能够走到这一步的步骤，找到最远的步骤，再次尝试走到这一步最远的步骤

		尝试倒推：
	*/
	point := len(nums) - 1 // 位置在最后
	steps := 0             // 初始化步骤
	for point > 0 {
		for i := 0; i < point; i++ {
			if i+nums[i] >= point {
				point = i
				steps++
			}
		}
	}
	return steps
}

func main() {
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
}
