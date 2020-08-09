package main

import "fmt"

//给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
//
//示例 1:
//
//输入: 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1。
//示例 2:
//
//输入: 10
//输出: 36
//解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
//说明: 你可以假设 n 不小于 2 且不大于 58。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/integer-break
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(integerBreak(10))
}

func integerBreak(n int) int {
	data := make(map[int]int)
	return integerBreak_(n, data)
}

func integerBreak_(n int, data map[int]int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	if value, ok := data[n]; ok {
		return value
	}
	res := -1
	for i := 1; i <= n - 1; i++ {
		res = max3(res, i * (n-i), i * integerBreak_(n-i, data))
	}
	data[n] = res
	return res
}

func max3(a, b, c int) int {
	temp := a
	if a < b {
		temp = b
	}
	if temp > c {
		return temp
	}
	return c
}