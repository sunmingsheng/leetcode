package main

import "fmt"

//给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。
//
//你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
//
//返回获得利润的最大值。
//
//注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
//
//示例 1:
//
//输入: prices = [1, 3, 2, 8, 4, 9], fee = 2
//输出: 8
//解释: 能够达到的最大利润:
//在此处买入 prices[0] = 1
//在此处卖出 prices[3] = 8
//在此处买入 prices[4] = 4
//在此处卖出 prices[5] = 9
//总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
}

func maxProfit(prices []int, fee int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}
	cash := 0
	hold := 0 - prices[0]
	for i := 1; i < length; i++ {
		cash = max(cash, hold + prices[i] - fee)
		hold = max(hold, cash - prices[i])
	}
	return cash
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func maxProfit(prices []int, fee int) int {
//	m := make(map[int]int)
//	return maxProfit_m(prices, fee, 0, len(prices) - 1, m)
//}
//
//func maxProfit_m(prices []int, fee , l , r int, m map[int]int) int {
//	length := r - l + 1
//	if length <= 1 {
//		return 0
//	}
//	if value, ok := m[l]; ok {
//		return value
//	}
//	res := 0
//	for i := l; i < r; i++ {
//		for j := i + 1; j < r + 1; j ++ {
//			temp := prices[j] - prices[i] - fee + maxProfit_m(prices, fee, j + 1, r, m)
//			if temp > res {
//				res = temp
//			}
//		}
//	}
//	m[l] = res
//	return res
//}
