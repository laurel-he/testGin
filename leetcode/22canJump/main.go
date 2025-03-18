package main

import "fmt"

func canJump(nums []int) bool {
	/*
		最远能够到达的位置=当前所在位置+当前所在位置的值
		如果最远能够到达位置比当前所在位置小，那么不能到达，否则都能到达
	*/
	canReach := 0
	for i := 0; i < len(nums); i++ {
		if i > canReach {
			return false
		}
		if i+nums[i] > canReach {
			canReach = i + nums[i]
		}
	}
	return true
}

func main() {
	res1 := canJump([]int{2, 3, 1, 1, 4})
	res2 := canJump([]int{3, 2, 1, 0, 4})
	fmt.Println(res1)
	fmt.Println(res2)
}
