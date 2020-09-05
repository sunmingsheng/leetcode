package main

//输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
//
// 
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
//
//提示：
//
//节点总数 <= 10000
//注意：本题与主站 113 题相同：h
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	dfs(root, 0, sum, []int{}, &res)
	return res
}

func dfs(node *TreeNode, count int, target int, data []int, res *[][]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		if count + node.Val == target {
			*res = append(*res, append(data, node.Val))
		}
		return
	}
	temp := make([]int, len(data))
	copy(temp, data)
	temp = append(temp, node.Val)
	if node.Left != nil {
		dfs(node.Left, count+node.Val, target, temp, res)
	}
	if node.Right != nil {
		dfs(node.Right, count+node.Val, target, temp, res)
	}
}