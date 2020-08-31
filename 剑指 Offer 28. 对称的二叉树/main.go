package main

//请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
//
//例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3
//但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3
//
// 
//
//示例 1：
//
//输入：root = [1,2,2,3,4,4,3]
//输出：true
//示例 2：
//
//输入：root = [1,2,2,null,3,null,3]
//输出：false
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return check(left.Left, right.Right) && check(left.Right, right.Left)
}

