package main

import "fmt"

//给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
//
//示例 1:
//
//输入: "(()"
//输出: 2
//解释: 最长有效括号子串为 "()"
//示例 2:
//
//输入: ")()())"
//输出: 4
//解释: 最长有效括号子串为 "()()"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-valid-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(longestValidParentheses("(()"))
}

func longestValidParentheses(s string) int {
	length := len(s)
	if length <= 1 {
		return length
	}
	dp := make([]int,length)
	//()(()) i = 2 j = 5
	//dp[i] = dp[i-1] + 2 + 上个元素的dp值
	//i - dp[i-1] - 1
	//i - dp[i-1] - 2
	max := 0
	for i := 1; i < length; i++ {
		if s[i] == ')' {
			if s[i - dp[i-1] - 1] == '(' {
				dp[i] = dp[i-1] + 2
				if i - dp[i-1] - 2 >= 0 {
					dp[i] += dp[i - dp[i-1] - 2]
				}
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}
	return max
}


