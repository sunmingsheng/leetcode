package main

import "fmt"

//给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
//
//示例 1:
//
//输入: "aba"
//输出: True
//示例 2:
//
//输入: "abca"
//输出: True
//解释: 你可以删除c字符。
//注意:
//
//字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-palindrome-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(validPalindrome("abcaa"))
}

func validPalindrome(s string) bool {
	length := len(s)
	if length <= 1 {
		return true
	}
	i, j := 0, length - 1
	for i < j {
		if s[i:i+1] == s[j:j+1] {
			i++
			j--
			continue
		} else {
			return helper(s[:i] + s[i+1:]) || helper(s[:j] + s[j+1:])
		}
	}
	return true
}

func helper(s string) bool {
	length := len(s)
	if length <= 1 {
		return true
	}
	i, j := 0, length - 1
	for i < j {
		if s[i:i+1] == s[j:j+1] {
			i++
			j--
			continue
		} else {
			return false
		}
	}
	return true
}
