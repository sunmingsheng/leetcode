package main

import "fmt"

//给你一棵以 root 为根的二叉树和一个整数 target ，请你删除所有值为 target 的 叶子节点 。
//
//注意，一旦删除值为 target 的叶子节点，它的父节点就可能变成叶子节点；如果新叶子节点的值恰好也是 target ，那么这个节点也应该被删除。
//
//也就是说，你需要重复此过程直到不能继续删除。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/delete-leaves-with-a-given-value
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1}
	fmt.Println(removeLeafNodes(root, 1))
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	for {
		has := false
		root = remove(root, target, &has)
		if !has {
			break
		}
	}
	return root
}

func remove(root *TreeNode, target int, has *bool) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil && root.Val == target {
		*has = true
		return nil
	}
	root.Right = remove(root.Right, target, has)
	root.Left = remove(root.Left, target, has)
	return root
}



