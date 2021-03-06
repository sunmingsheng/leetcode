package main

import (
	"fmt"
	"math"
)

//给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。
//
//示例：
//
//输入: S = "ADOBECODEBANC", T = "ABC"
//输出: "BANC"
//说明：
//
//如果 S 中不存这样的子串，则返回空字符串 ""。
//如果 S 中存在这样的子串，我们保证它是唯一的答案。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-window-substring
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	s1 := "bbaa"
	s2 := "aba"
	fmt.Println(minWindow(s1, s2))
}

func minWindow(s string, t string) string {
	need, window := map[byte]int{}, map[byte]int{}
	for i := range t {
		need[t[i]]++
	}
	left, right := 0, 0
	valid := 0
	index, length := 0, math.MaxInt32
	for right < len(s) {
		tempAdd := s[right]
		right++

		if _, ok := need[tempAdd]; ok {
			window[tempAdd]++
			if window[tempAdd] == need[tempAdd] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				index = left
				length = right - left
			}
			tempDel := s[left]
			left++

			if _, ok := need[tempDel]; ok {
				window[tempDel]--
				if window[tempDel] < need[tempDel] {
					valid--
				}
			}
		}

	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[index : index+length]
}

//func getM(t string) map[string]int {
//	m := make(map[string]int)
//	for i := 0; i < len(t); i++ {
//		if _, ok := m[t[i:i+1]]; ok {
//			m[t[i:i+1]] += 1
//		} else {
//			m[t[i:i+1]] = 1
//		}
//	}
//	return m
//}
//
//func minWindow(s string, t string) string {
//	sLen := len(s)
//	tLen := len(t)
//	res  := ""
//	bakMap := getM(t)
//	j := 0
//	for i := 0 ; i < sLen - tLen + 1; i++ {
//		if i + tLen - 1 > j {
//			j = i + tLen - 1
//		}
//		for ; j < sLen; j++ {
//			m := make(map[string]int)
//			for z := i; z <= j; z++ {
//				str := s[z:z+1]
//				if _, ok := m[str]; ok {
//					m[str] += 1
//				} else {
//					m[str] = 1
//				}
//			}
//			flag := true
//			for str, c1 := range bakMap {
//				if c2, ok := m[str]; !ok || c1 > c2 {
//					flag = false
//					break
//				}
//			}
//			if flag {
//				if res == "" || (res != "" && j - i + 1 < len(res)) {
//					res = s[i:j+1]
//				}
//				break
//			}
//		}
//	}
//	return res
//}