package main

import "fmt"

//请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
//
// 
//
//示例 1:
//
//输入: "abcbbcbb"
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
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length <= 1 {
		return length
	}
	maxLength, i, j, m := 1, 0, 0, make(map[string]struct{})
	for ; i < length; i++ {
		if _, ok := m[s[i:i+1]]; !ok {
			m[s[i:i+1]] = struct{}{}
			if i - j + 1 > maxLength {
				maxLength = i - j + 1
			}
		} else {
			for {
				if _, ok := m[s[i:i+1]]; !ok {
					break
				}
				delete(m, s[j:j+1])
				j++
			}
			m[s[i:i+1]] = struct{}{}
		}
	}
	return maxLength
}