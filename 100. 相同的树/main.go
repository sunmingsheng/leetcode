package main

import "fmt"

//给定两个二叉树，编写一个函数来检验它们是否相同。
//
//如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
//
//示例 1:
//
//输入:       1         1
/// \       / \
//2   3     2   3
//
//[1,2,3],   [1,2,3]
//
//输出: true
//示例 2:
//
//输入:      1          1
///           \
//2             2
//
//[1,2],     [1,null,2]
//
//输出: false
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/same-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	p := &TreeNode{Val:1}
	q := &TreeNode{Val:1}
	fmt.Println(isSameTree(p, q))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	if !isSameTree(p.Right, q.Right) || !isSameTree(p.Left, q.Left) {
		return false
	}
	return true
}
