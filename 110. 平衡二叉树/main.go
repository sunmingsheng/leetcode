package main

import "fmt"

//给定一个二叉树，判断它是否是高度平衡的二叉树。
//
//本题中，一棵高度平衡二叉树定义为：
//
//一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
//
//示例 1:
//
//给定二叉树 [3,9,20,null,null,15,7]
//
//3
/// \
//9  20
///  \
//15   7
//返回 true 。
//
//示例 2:
//
//给定二叉树 [1,2,2,3,3,null,null,4,4]
//
//1
/// \
//2   2
/// \
//3   3
/// \
//4   4
//返回 false 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/balanced-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(isBalanced(&TreeNode{Val:1}))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lHeight, rHeight := 0, 0
	getHeight(root.Left, 1, &lHeight)
	getHeight(root.Right, 1, &rHeight)

	if lHeight > rHeight && lHeight - rHeight > 1 {
		return false
	}

	if lHeight < rHeight && lHeight - rHeight < -1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getHeight(node *TreeNode, level int, maxHeight *int)  {
	if node == nil {
		return
	}
	if level > *maxHeight {
		*maxHeight = level
	}
	getHeight(node.Left, level + 1, maxHeight)
	getHeight(node.Right, level + 1, maxHeight)
}
