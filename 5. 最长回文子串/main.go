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
	fmt.Println(longestPalindrome("abacab"))
}

func longestPalindrome(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}
	res := s[:1]
	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j -- {
			strLen := j - i + 1
			if strLen <= len(res) {
				continue
			}
			mid  := i + (j - i) / 2
			strA := ""
			strB := ""
			if (j - i) % 2 == 0 { //奇数
				strA = s[i:mid]
				strB = s[mid+1:j+1]
			}  else { //偶数
				strA = s[i:mid+1]
				strB = s[mid+1:j+1]
			}
			//fmt.Println(i, mid, j, strA, strB)
			flag := true
			for z := 0 ; z < len(strA); z ++ {
				if strA[z] != strB[len(strA) - z - 1] {
					flag = false
					break
				}
			}
			if flag {
				res = s[i : j+1]
			}
		}
	}
	return res
}
