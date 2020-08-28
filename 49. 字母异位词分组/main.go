package main

import "fmt"

//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
//示例:
//
//输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//["ate","eat","tea"],
//["nat","tan"],
//["bat"]
//]
//说明：
//
//所有输入均为小写字母。
//不考虑答案输出的顺序。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/group-anagrams
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	length := len(strs)
	if length == 0 {
		return [][]string{}
	}
	if length == 1 {
		return [][]string{strs}
	}
	res := [][]string{}
	m := make(map[[26]int][]string)
	for i := 0; i < length; i++ {
		data := [26]int{}
		for _, value := range strs[i] {
			data[int(value) - int('a')] += 1
		}
		if _, ok := m[data]; !ok {
			m[data] = []string{}
		}
		m[data] = append(m[data], strs[i])
	}
	for _, value := range m {
		res = append(res, value)
	}
	return res
}