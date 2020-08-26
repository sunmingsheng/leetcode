package main

import "fmt"

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
//示例 1：
//
//输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//示例 2：
//
//输入: "cbbd"
//输出: "bb"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-palindromic-substring
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	//两种模式 一种是aba 一种是baab
	fmt.Println(longestPalindrome("aaaa"))
}

func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	start := 0
	maxLength := 1
	for i := 0; i < len(s); i++ {
		helper(s, i - 1, i + 1, &start, &maxLength)
		helper(s, i, i + 1, &start, &maxLength)
	}
	return s[start:start + maxLength]
}

func helper(s string, left , right int, start, maxLength *int) {
	for {
		if left < 0 || right >= len(s) {
			return
		}
		if s[left] != s[right] {
			return
		}
		if right - left + 1 > *maxLength {
			*maxLength = right - left + 1
			*start = left
		}
		left--
		right++
	}
}
