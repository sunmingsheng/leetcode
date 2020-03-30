package main

import "fmt"

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
//假设一个二叉搜索树具有如下特征：
//
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
//示例 1:
//
//输入:
//2
/// \
//1   3
//输出: true
//示例 2:
//
//输入:
//5
/// \
//1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/validate-binary-search-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//      10
//5           15
//          6    20
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//[10,5,15,null,null,6,20]
func main() {
	root := &TreeNode{Val:1}
	fmt.Println(isValidBST(root))
}

func isValidBST(root *TreeNode) bool {
	s := &[]int{}
	orderTraversal(root, s)
	length := len(*s)
	flag := true
	for key, value := range *s {
		if key <= length && value >= (*s)[key + 1] {
			flag = false
			break
		}
	}
	return flag
}

func orderTraversal(node *TreeNode, s *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		orderTraversal(node.Left, s)
	}
	*s = append(*s, node.Val)
	if node.Right != nil {
		orderTraversal(node.Right, s)
	}
}
