package main

import "fmt"

//根据一棵树的中序遍历与后序遍历构造二叉树。
//
//注意:
//你可以假设树中没有重复的元素。
//
//例如，给出
//
//中序遍历 inorder = [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3]
//返回如下的二叉树：
//
//3
/// \
//9  20
///  \
//15   7
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	inorder := []int{2,3,1}
	postorder := []int{3,2,1}
	// 1
	//
	fmt.Println(buildTree(inorder, postorder))
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	l := len(postorder)-1
	for k := range inorder {
		if inorder[k] == postorder[l] {
			return &TreeNode{
				Val:   inorder[k],
				Left:  buildTree(inorder[:k], postorder[:k]),
				Right: buildTree(inorder[k+1:], postorder[k:l]),
			}
		}
	}
	return nil
}

