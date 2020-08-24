package main

import "fmt"

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//示例:
//
//输入："23"
//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//说明:
//尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(letterCombinations("20"))
}

func letterCombinations(digits string) []string {
	m := make(map[string][]string)
	m["2"] = []string{"a","b","c"}
	m["3"] = []string{"d","e","f"}
	m["4"] = []string{"g","h","i"}
	m["5"] = []string{"j","k","l"}
	m["6"] = []string{"m","n","o"}
	m["7"] = []string{"p","q","r", "s"}
	m["8"] = []string{"t","u","v"}
	m["9"] = []string{"w","x","y", "z"}
	length := len(digits)
	if length == 0 {
		return []string{}
	}
	if length == 1 {
		if results, ok := m[digits]; ok {
			return results
		}
		return []string{}
	}
	results := []string{}
	for i := 0; i < length; i++ {
		if _, ok := m[digits[i:i+1]]; !ok {
			continue
		}
		if len(results) == 0 {
			results = m[digits[i:i+1]]
		} else {
			temp := []string{}
			for _, out := range results {
				for _, inner := range m[digits[i:i+1]] {
					temp = append(temp, out + inner)
				}
			}
			results = temp
		}
	}
	return results
}
