package main

import "fmt"

//给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例:
//给定如下二叉树，以及目标和 sum = 22，
//
//5
/// \
//4   8
///   / \
//11  13  4
///  \    / \
//7    2  5   1
//返回:
//
//[
//[5,4,11,2],
//[5,8,4,5]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/path-sum-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(pathSum(&TreeNode{Val:1},1))
	num1 := []int{1}
	num2 := num1
	num2 = append(num2, 2)
	num1 = append(num1, 3)
	fmt.Println(num1)
	fmt.Println(num2)
}

func test(nums []int) {
	nums = append(nums, 23)
	test(nums)
}

func pathSum(root *TreeNode, sum int) [][]int {
	results := &[][]int{}
	if root == nil {
		return *results
	}
	dfs(root, sum, []int{}, results)
	return *results
}

func dfs(node *TreeNode, sum int, nodeValue []int, results *[][]int) {
	lists := make([]int, len(nodeValue))
	copy(lists, nodeValue) //必须要用copy函数复制一个切片 []int传来是一个地址 append修改后，之前的切片len不变，后面继续append,之前写入的数据会被覆盖
	if node.Left == nil && node.Right == nil {
		if Sum(nodeValue) + node.Val == sum {
			lists = append(lists, node.Val)
			*results = append(*results, lists)
		}
		return
	}
	lists = append(lists, node.Val)
	if node.Left != nil {
		dfs(node.Left, sum, lists, results)
	}
	if node.Right != nil {
		dfs(node.Right, sum, lists, results)
	}
}

func Sum(nums []int) int {
	data := 0
	for _, value := range nums {
		data += value
	}
	return data
}