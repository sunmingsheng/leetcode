package main

import (
	"fmt"
	"strings"
)

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