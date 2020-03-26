package main

import "fmt"

//给定二叉搜索树的根结点 root，返回 L 和 R（含）之间的所有结点的值的和。
//
//二叉搜索树保证具有唯一的值。
//
// 
//
//示例 1：
//
//输入：root = [10,5,15,3,7,null,18], L = 7, R = 15
//输出：32
//示例 2：
//
//输入：root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
//输出：23
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/range-sum-of-bst
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1}
	fmt.Println(rangeSumBST(root, 0, 2))
}

//中序遍历获取排序好的数据
func rangeSumBST(root *TreeNode, l int, r int) int {
	if root == nil || l > r {
		return 0
	}
	s := &[]int{}
	middleTraversal(root, s, l, r)
	results := 0
	for _, value := range *s {
		results += value
	}
	return results
}

func middleTraversal(node *TreeNode, s *[]int, l , r int) {
	if node == nil {
		return
	}
	if node.Val >= l {
		middleTraversal(node.Left, s, l, r)
	}
	if node.Val >= l && node.Val <= r {
		*s = append(*s, node.Val)
	}
	if node.Val <= r {
		middleTraversal(node.Right, s, l, r)
	}
}


