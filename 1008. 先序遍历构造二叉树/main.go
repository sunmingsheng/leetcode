package main

import (
	"fmt"
)

//返回与给定先序遍历 preorder 相匹配的二叉搜索树（binary search tree）的根结点。
//
//(回想一下，二叉搜索树是二叉树的一种，其每个节点都满足以下规则，对于 node.left 的任何后代，值总 < node.val，而 node.right 的任何后代，值总 > node.val。此外，先序遍历首先显示节点的值，然后遍历 node.left，接着遍历 node.right。）
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(bstFromPreorder([]int{1,2,3}))
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	head := &TreeNode{Val:preorder[0]}
	for _, value := range preorder[1:] {
		addNodeToTree(head, value)
	}
	return head
}

func addNodeToTree(head *TreeNode, val int) {
	if head.Val >= val {
		if head.Left == nil {
			head.Left = &TreeNode{Val:val}
		} else {
			addNodeToTree(head.Left, val)
		}
	} else {
		if head.Right == nil {
			head.Right = &TreeNode{Val:val}
		} else {
			addNodeToTree(head.Right, val)
		}
	}
}


