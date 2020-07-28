package main

import "fmt"

//给你一棵根为 root 的二叉树，请你返回二叉树中好节点的数目。
//
//「好节点」X 定义为：从根到该节点 X 所经过的节点中，没有任何节点的值大于 X 的值。
//
// 
//
//示例 1：
//
//
//
//输入：root = [3,1,4,3,null,1,5]
//输出：4
//解释：图中蓝色节点为好节点。
//根节点 (3) 永远是个好节点。
//节点 4 -> (3,4) 是路径中的最大值。
//节点 5 -> (3,4,5) 是路径中的最大值。
//节点 3 -> (3,1,3) 是路径中的最大值。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/count-good-nodes-in-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1}
	fmt.Println(goodNodes(root))
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	num := 0
	travel(root, &num, root.Val)
	return num
}

func travel(node *TreeNode, num *int, maxValue int) {
	if node == nil {
		return
	}
	if node.Val >= maxValue {
		*num = *num + 1
		maxValue = node.Val
	}
	travel(node.Left, num, maxValue)
	travel(node.Right, num, maxValue)
}