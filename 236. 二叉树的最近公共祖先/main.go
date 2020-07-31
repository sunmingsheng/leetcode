package main

import "fmt"

//给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
//百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
//
//例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]
// 
//
//示例 1:
//
//输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
//输出: 3
//解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
//示例 2:
//
//输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
//输出: 5
//解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
// 
//
//说明:
//
//所有节点的值都是唯一的。
//p、q 为不同节点且均存在于给定的二叉树中。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	left := &TreeNode{Val:2}
	right := &TreeNode{Val:3}
	root := &TreeNode{Val:1, Left:left, Right:right}
	fmt.Println(lowestCommonAncestor(root, left, right))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	nodes1 := &[]*TreeNode{}
	tempNodes1 := []*TreeNode{}
	nodes2 := &[]*TreeNode{}
	tempNodes2 := []*TreeNode{}
	travel(root, p.Val, tempNodes1, nodes1)
	travel(root, q.Val, tempNodes2, nodes2)
	for i := len(*nodes1) - 1; i >= 0; i-- {
		for _, value := range *nodes2 {
			if (*nodes1)[i] == value {
				return value
			}
		}
	}
	return nil
}

func travel(node *TreeNode, val int, temp []*TreeNode, nodes *[]*TreeNode) {
	if node == nil {
		return
	}
	//data := make([]*TreeNode, len(temp))
	//copy(data, temp)
	//data = append(data, node)
	temp = append(temp, node)
	if node.Val == val {
		nodes = &temp
		//*nodes = data
		return
	}
	travel(node.Left, val, temp, nodes)
	travel(node.Right, val, temp, nodes)
}