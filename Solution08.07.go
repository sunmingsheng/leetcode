package main

import (
	"fmt"
)

//无重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合，字符串每个字符均不相同。
//
//示例1:
//
//输入：S = "qwe"
//输出：["qwe", "qew", "wqe", "weq", "ewq", "eqw"]
//示例2:
//
//输入：S = "ab"
//输出：["ab", "ba"]
//提示:
//
//字符都是英文字母。
//字符串长度在[1, 9]之间。

func main() {
	fmt.Println(permutation("qwe"))
}

//解题思路
//回溯法记录路径
//back(runes []rune, path string, results *[]string)
//runes记录的是未处理的字符数据
//path记录的是已经处理的数据
func permutation(str string) []string {
	results := []string{}
	runes := []rune(str)
	path := ""
	back(runes, path, &results)
	return results
}

func back(runes []rune, path string, results *[]string) {
	if len(runes) == 0 {
		*results = append(*results, path)
		return
	}
	for key, value := range runes {
		temp := make([]rune, key)
		copy(temp, runes[:key])
		temp = append(temp, runes[key+1:]...)
		back(temp, path + string(value), results)
	}
}

