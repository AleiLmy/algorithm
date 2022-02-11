/**
 * @Author: xuelei02
 * @Description:
 * @File:  stock
 * @Version: 1.0.0
 * @Date: 2021/10/28 19:06
 */

package main

import (
	"fmt"
	"math"
)

// 买卖股票的最佳时机

// MaxStock 只能买卖一次
func MaxStock(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}

	// 初始化
	dp[0][0] = -prices[0] // 第一天买入
	dp[0][1] = 0          // 第一天卖出

	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0])
	}
	return dp[length-1][1]
}

func MaxStock2(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	var (
		ret = 0
		min = math.MaxInt
	)
	for _, v := range prices {
		if v < min {
			min = v
		}
		if v-min > ret {
			ret = v - min
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// fmt.Println(MaxStock1([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(MaxStock([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(MaxStock2([]int{7, 1, 5, 3, 6, 4}))
}

// MaxStock1 可以多次买卖
// dp[i][0] 表示第i天没有股票的最大收益
// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+price[i]) // 前一天有 今天卖了
// dp[i][1] 第i天有股票的最大收益
// dp[i][1] = max(dp[i-1][1],dp[i-1][0]-price[i])
func MaxStock1(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[length-1][0]
}
