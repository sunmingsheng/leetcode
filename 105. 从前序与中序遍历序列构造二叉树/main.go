package main

import "fmt"

//根据一棵树的前序遍历与中序遍历构造二叉树。
//
//注意:
//你可以假设树中没有重复的元素。
//
//例如，给出
//
//前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//返回如下的二叉树：
//
//3
/// \
//9  20
///  \
//15   7
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(buildTree([]int{1,2,3}, []int{2,3,1}))
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	val := preorder[0]
	for key, value := range inorder {
		if val == value {
			return &TreeNode{
				Val:   val,
				Left:  buildTree(preorder[1:key + 1], inorder[:key]),
				Right: buildTree(preorder[key + 1:], inorder[key+1:]),
			}
		}
	}
	return nil
}
