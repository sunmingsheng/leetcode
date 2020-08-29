package main

//输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。
//
//例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。
//
// 
//
//示例 1：
//
//输入：n = 121 10 11 12
//输出：5
//示例 2：
//
//输入：n = 13
//输出：6
// 
//
//限制：
//
//1 <= n < 2^31
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


func countDigitOne(n int) int {
	digitNum, sum := 1, 0
	high, cur, low := n/10, n%10, 0
	for high != 0 || cur != 0 {
		// 固定位 计算数量
		if cur < 1 {
			sum += high * digitNum
		} else if cur == 1 {
			sum += high*digitNum + low + 1
		} else {
			sum += (high + 1) * digitNum
		}
		// 换下一位 更新高低位 及digit数量级
		low = low + cur*digitNum
		high, cur = high/10, high%10    //这里 cur 是等于 前一个 high 模10
		digitNum = digitNum * 10
	}
	return sum
}

