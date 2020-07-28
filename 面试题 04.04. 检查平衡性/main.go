package main

import "fmt"

//实现一个函数，检查二叉树是否平衡。在这个问题中，平衡树的定义如下：任意一个节点，其两棵子树的高度差不超过 1。
//
//
//示例 1:
//给定二叉树 [3,9,20,null,null,15,7]
//3
/// \
//9  20
///  \
//15   7
//返回 true 。
//示例 2:
//给定二叉树 [1,2,2,3,3,null,null,4,4]
//1
/// \
//2   2
/// \
//3   3
/// \
//4   4
//返回 false
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/check-balance-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
 	Val int
 	Left *TreeNode
 	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(isBalanced(root))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lHeight := 0
	rHeight := 0
	getHeight(root.Left, 0, &lHeight)
	getHeight(root.Right, 0, &rHeight)

	if lHeight > rHeight && (lHeight - rHeight) > 1 {
		return false
	}
	if lHeight < rHeight && (lHeight - rHeight) < -1 {
		return false
	}
	return isBalanced(root.Right) && isBalanced(root.Left)
}

func getHeight(node *TreeNode, height int, maxHeight *int) {
	if node == nil {
		return
	}
	height += 1
	if height > *maxHeight {
		*maxHeight += 1
	}
	getHeight(node.Left, height, maxHeight)
	getHeight(node.Right, height, maxHeight)
}