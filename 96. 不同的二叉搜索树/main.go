package main

import "fmt"

//给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
//
//示例:
//
//输入: 3
//输出: 5
//解释:
//给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
//
//1         3     3      2      1
//\       /     /      / \      \
//3     2     1      1   3      2
///     /       \                 \
//2     1         2                 3
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/unique-binary-search-trees
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(numTrees(3))
}

func numTrees(n int) int {
	s := []int{}
	for i := 1; i <= n; i++ {
		s = append(s, i)
	}
	result := 0
	for i := 0; i < n; i++ {
		leftCount := 1
		rightCount := 1
		order(s[i], s[0:i], s[i+1:], &leftCount)
		order(s[i], s[0:i], s[i+1:], &rightCount)
		result += rightCount * leftCount
	}
	return result
}

func order(middle int, left []int, right []int, count *int) {
	if len(left) == 0 && len(right) == 0 {
		*count = *count + 1
		return
	}

	for key, value := range left {

	}
}
