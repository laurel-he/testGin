package main

// rotate 轮转数组：所有数据往后移动k位，超出的往前补充
func rotate(nums []int, k int) {
	/*
		考虑轮转的话，第n个元素的位置是原本的第几个元素？
		第n个元素要往后调k个，那么位置应该是n+k，当n+k>len时，实际上就是取余
		还要注意数据往后调了，那么对应坑位的值可能有变化，需要单独存储一份
	*/
	originNums := make([]int, len(nums))
	copy(originNums, nums)
	for i := 0; i < len(nums); i++ {
		if i+k == len(nums)-1 {
			nums[len(nums)-1] = originNums[i]
			continue
		}
		nums[(i+k)%len(nums)] = originNums[i]
	}
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
