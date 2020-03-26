package main

import (
	"fmt"
	"strings"
)

//给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
//
//例如，给出 n = 3，生成结果为：
//
//[
//"((()))",
//"(()())",
//"(())()",
//"()(())",
//"()()()"
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/generate-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	results := generateParenthesis(3)
	for _, value := range results {
		fmt.Println(value)
	}
}

func generateParenthesis(n int) []string {
	results := []string{}
	leftNum := n - 1
	rightNum := n
	path := "("
    back3(leftNum, rightNum, path, &results)
	return results
}

func back3(leftNum, rightNum int, path string, results *[]string) {
	if leftNum > rightNum {
		return
	}
	if leftNum <= 0 || rightNum <= 0 {
		if leftNum <= 0 && rightNum > 0 {
			path += strings.Repeat(")", rightNum)
		}
		if leftNum > 0 && rightNum <= 0{
			path += strings.Repeat(")", leftNum)
		}
		*results = append(*results, path)
		return
	}
	back3(leftNum - 1, rightNum, path + "(", results)
	back3(leftNum, rightNum - 1, path + ")", results)
}