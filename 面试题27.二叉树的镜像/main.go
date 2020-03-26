package main

import "fmt"

//请完成一个函数，输入一个二叉树，该函数输出它的镜像。
//
//例如输入：
//
//     4
//   /   \
//  2     7
// / \   / \
//1   3 6   9
//镜像输出：
//
//     4
//   /   \
//  7     2
// / \   / \
//9   6 3   1
//
// 
//
//示例 1：
//
//输入：root = [4,2,7,1,3,6,9]
//输出：[4,7,2,9,6,3,1]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:4}
	fmt.Println(mirrorTree(root))
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
    mirror(root)
	return root
}

func mirror(node *TreeNode) {
	if node == nil {
		return
	}
	temp := node.Right
	node.Right = node.Left
	node.Left = temp
	mirror(node.Right)
	mirror(node.Left)
}