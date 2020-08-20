package main

import (
	"fmt"
)

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb" i = 3 j = 0 z = 2
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

func main() {
	//abcdbcbb
	fmt.Println(lengthOfLongestSubstring("aab"))
}

//窗口滑动法
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length <= 1 {
		return  length
	}
	maxLength := 0
	i,j := 0,0
	set := []string{}
	for ;i < length;i++ {
		temp := s[i:i+1]
		flag := checkHas(set, temp)
		if !flag {
			set = append(set, temp)
			if maxLength < len(set) {
				maxLength = len(set)
			}
		} else {
			for checkHas(set, temp) {
				set = set[1:]
				j++
			}
			set = append(set, temp)
			if maxLength < len(set) {
				maxLength = len(set)
			}
		}
	}
	return maxLength
}

func checkHas(set []string, target string) bool {
	for _, value := range set {
		if value == target {
			return true
		}
	}
	return false
}

