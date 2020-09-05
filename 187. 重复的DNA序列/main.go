package main

import "fmt"

//所有 DNA 都由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：“ACGAATTCCG”。在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。
//
//编写一个函数来查找目标子串，目标子串的长度为 10，且在 DNA 字符串 s 中出现次数超过一次。
//
// 
//
//示例：
//
//输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
//输出：["AAAAACCCCC", "CCCCCAAAAA"]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/repeated-dna-sequences
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))
}

func findRepeatedDnaSequences(s string) []string {
	length := len(s)
	res := []string{}
	if length <= 10 {
		return res
	}
	m := make(map[string]int)
	for i := 0; i <= length - 10; i++ {
		tmp := s[i:i+10]
		m[tmp]++
	}
	for key, value := range m {
		if value > 1 {
			res = append(res, key)
		}
	}
	return res
}
