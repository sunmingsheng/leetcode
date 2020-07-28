package main

import (
	"fmt"
)

//输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
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
// 
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//[0,2,4,1,null,3,-1,5,1,null,6,null,8]
//            0
//          2    4
//       1      3  -1
//     5  1      6   8
func main() {
	root := TreeNode{Val:0,
		Left:&TreeNode{
			Val:2,
			Left:  nil,
			Right: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val:1},
				Right: &TreeNode{Val:5},
			},
		},
		Right:&TreeNode{
			Val:   4,
			Left:  &TreeNode{
				Val:   3,
				Left:  nil,
				Right: &TreeNode{Val:6},
			},
			Right: &TreeNode{
				Val:   -1,
				Left:  nil,
				Right: &TreeNode{Val:8},
			},
		},
	}
	fmt.Println(isBalanced(&root))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftH := 0
	rightH := 0
	getHeight(root.Left, 0,  &leftH)
	getHeight(root.Right, 0, &rightH)
	if leftH - rightH > 0 && leftH - rightH > 1 {
		return false
	}
	if leftH - rightH < 0 && rightH - leftH > 1{
		return false
	}
	return isBalanced(root.Right) && isBalanced(root.Left)
}

func getHeight(node *TreeNode, level int, height *int) {
	if node == nil {
		return
	}
	level += 1
	if level > *height {
		*height = level
	}
	getHeight(node.Right, level, height)
	getHeight(node.Left, level, height)
}
