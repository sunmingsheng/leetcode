package main

import "fmt"

//给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
//
//示例 1:
//
//输入: "abab"
//
//输出: True
//
//解释: 可由子字符串 "ab" 重复两次构成。
//示例 2:
//
//输入: "aba"
//
//输出: False
//示例 3:
//
//输入: "abcabcabcabc"
//
//输出: True
//
//解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/repeated-substring-pattern
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(repeatedSubstringPattern("abab"))
}

func repeatedSubstringPattern(s string) bool {
	length := len(s)
	if length <= 1 {
		return false
	}
	for i := 2; i <= length; i++ {
		length1 := length / i
		if length1 * i != length {
			continue
		}
		flag := true
		for j := 0; j < i - 1; j++ {
			if s[j*length1:(j+1)*length1] != s[(j+1)*length1:(j+2)*length1] {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}