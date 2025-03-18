package main

import "fmt"

// maxProfit 最大利润：允许反复买进卖出，但是同时只能持有一支股票
func maxProfit(prices []int) int {
	/*
		每天有两种选择：持有股票/不持有股票，直到卖出才能计算利润
		（1）持有股票(今天不会卖出)：与前一天有关，可能（前一天也持有股票）/（前一天没有股票，今天才买入）
			（1.1）前一天也持有股票：那么最大利润就是前一天持有股票的利润，因为今天不会卖出
			（1.2）前一天没有股票，今天才买入：最大利润为前一天没有股票的利润-今天的买入价
		（2）没有股票（今天不会买入）：与前一天有关，可能（前一天也没有股票）/（前一天有股票但是今日卖出）
			（2.1）前一天也没有股票：那么最大利润就是前一天没有股票情况的利润
			（2.2）前一天有股票但是今日卖出：最大利润为前一天有股票的利润+今天的售出价
		假设没有股票为dp[i][0],持有股票为dp[i][1]
		那么 dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		得到以上状态转移方程：
		最后一天得到最大利润，最后一天手上一定没有股票，否则利润会降低，也就是dp[n-1][0]就是最大利润
	*/
	var dp [][]int
	dp = make([][]int, len(prices))
	dp0, dp1 := 0, 0
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			dp[0] = make([]int, 2)
			dp0 = 0          // 没有股票，即使是当天买入再卖出，依旧无受益，不买更没有收益
			dp1 = -prices[0] // 买入股票，扣除买入价
			fmt.Println("__1__", dp0)
			fmt.Println("__2__", dp1)
		} else {
			/*
				直接使用动态规划解答
				dp[i] = make([]int, 2)
				dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
				dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
			*/
			dp[i] = make([]int, 2)
			newDp0 := max(dp0, dp1+prices[i])
			newDp1 := max(dp1, dp0-prices[i])
			dp0 = newDp0
			dp1 = newDp1
		}
	}
	return dp0
}

// max 计算最大值
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
