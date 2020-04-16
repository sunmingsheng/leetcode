package main

import "fmt"

//给定一个二叉树，计算整个树的坡度。
//
//一个树的节点的坡度定义即为，该节点左子树的结点之和和右子树结点之和的差的绝对值。空结点的的坡度是0。
//
//整个树的坡度就是其所有节点的坡度之和。
//
//示例:
//
//输入:
//1
///   \
//2     3
//输出: 1
//解释:
//结点的坡度 2 : 0
//结点的坡度 3 : 0
//结点的坡度 1 : |2-3| = 1
//树的坡度 : 0 + 0 + 1 = 1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-tilt
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1}
	fmt.Println(findTilt(root))
}


func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := findTilt(root.Left)
	right := findTilt(root.Right)
	if left > right {
		return left - right
	}
	return right - left
}

func print(node *TreeNode, result *int) int {
	if node == nil {
		return 0
	}
	left := node.Val + print(node.Left, result)
	right := node.Val + print(node.Right, result)
	results := 0
	if left > right {
		results =  left - right
	} else {
		results = right - left
	}
	*result += results
	return results
}


