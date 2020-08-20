package main

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
//示例 1:
//
//输入: s = "anagram", t = "nagaram"
//输出: true
//示例 2:
//
//输入: s = "rat", t = "car"
//输出: false
//说明:
//你可以假设字符串只包含小写字母。
//
//进阶:
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-anagram
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func isAnagram(s string, t string) bool {
	length1 := len(s)
	length2 := len(t)
	if length1 != length2 {
		return false
	}
	m1 := make(map[string]int)
	m2 := make(map[string]int)
	for i:=0; i < length1; i++ {
		s1 := s[i:i+1]
		t1 := t[i:i+1]
		m1[s1] += 1
		m2[t1] += 1
	}
	if len(m1) != len(m2) {
		return false
	}
	for k,v := range m1 {
		if count, ok := m2[k];!ok || count != v {
			return false
		}
	}
	return true
}