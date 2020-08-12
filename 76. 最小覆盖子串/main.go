package main

import (
	"fmt"
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
	res := ""
	l1 := len(s)
	l2 := len(t)
	if l1 == 0 || l2 == 0 || l1 < l2 {
		return ""
	}
	m := make(map[string]int)
	for i := 0; i < l2; i++ {
		temp := t[i:i+1]
		if _, ok := m[temp]; ok {
			m[temp] += 1
		} else {
			m[temp] = 1
		}
	}
	i, j := 0, 0
	for ; i < l1; i++ {
		if i <= j {
			continue
		}
		tempM := make(map[string]int)
		for z := j; z <= i; z++ {
			if i - z + 1 < l2 {
				continue
			}
			tempS := s[z:z+1]
			if _, ok := tempM[tempS]; ok {
				tempM[tempS] += 1
			} else {
				tempM[tempS] = 1
			}
		}
		flag := true
		for key, value := range tempM {
			if count, ok := m[key]; !ok || count < value {
				flag = false
				break
			}
		}
		if flag {
			//满足条件
			if res != "" && len(s[j:i+1]) > len(res) {
				res = s[j:i+1]
			}
		}
	}
	return res
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