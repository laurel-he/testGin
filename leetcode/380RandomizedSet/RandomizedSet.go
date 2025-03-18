package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	nums []int
	res  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.res[val]; ok {
		return false
	}
	numsKey := len(this.nums)
	this.nums = append(this.nums, val)
	this.res[val] = numsKey
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if id, ok := this.res[val]; ok {
		//fmt.Println("___1___", this.nums)
		//fmt.Println("___1_1__", id)
		//this.nums = append(this.nums[:id], this.nums[id:]...)
		//fmt.Println("___2___", this.nums)
		//this.res[this.nums[id]] = id
		//delete(this.res, val)
		//return true
		last := len(this.nums) - 1
		this.nums[id] = this.nums[last]
		this.res[this.nums[id]] = id
		this.nums = this.nums[:last]
		delete(this.res, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.nums) == 1 {
		return this.nums[0]
	} else {
		random := rand.Intn(len(this.nums))
		//fmt.Println("len(this.nums)", len(this.nums))
		//fmt.Println("random", random)
		return this.nums[random]
	}
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	//obj := Constructor()
	//val := 3
	//param1 := obj.Insert(val)
	//param2 := obj.Remove(val)
	//param3 := obj.GetRandom()
	//fmt.Println(param1, param2, param3)
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(2)
	fmt.Println("random", random)
}
