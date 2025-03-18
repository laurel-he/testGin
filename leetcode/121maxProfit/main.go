package main

import "fmt"

// maxProfit 获取最大利润
func maxProfit(prices []int) int {
	/*
		根据题目可以绘制一个简单的波形图，从而能够得到结论，前期一定要拿到一个最小的买入价，后期查最大利润
		一次循环，每次循环计算最大利润
	*/
	minPrice := prices[0]
	maxProfitRes := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		profit := price - minPrice
		if profit > maxProfitRes {
			maxProfitRes = profit
		}
	}
	return maxProfitRes
}

func main() {
	profit := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(profit)
}
