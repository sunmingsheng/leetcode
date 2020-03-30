package main

import (
	"fmt"
	"sort"
)

//给你 root1 和 root2 这两棵二叉搜索树。
//
//请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.
//
//输入：root1 = [2,1,4], root2 = [1,0,3]
//输出：[0,1,1,2,3,4]
//示例 2：
//
//输入：root1 = [0,-10,10], root2 = [5,1,7,0,2]
//输出：[-10,0,0,1,2,5,7,10]
//示例 3：
//
//输入：root1 = [], root2 = [5,1,7,0,2]
//输出：[0,1,2,5,7]
//示例 4：
//
//输入：root1 = [0,-10,10], root2 = []
//输出：[-10,0,10]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/all-elements-in-two-binary-search-trees
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	node1 := &TreeNode{Val:1}
	node2 := &TreeNode{Val:2}
	fmt.Println(getAllElements(node1, node2))
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	s1 := &[]int{}
	orderTraversal(root1, s1)
	s2 := &[]int{}
	orderTraversal(root2, s2)
	s := make([]int, len(*s1))
	copy(s, *s1)
	s = append(s, *s2...)
	sort.Ints(s)
	return s
}

//中序遍历
func orderTraversal(node *TreeNode, s *[]int) {
	if node == nil {
		return
	}
	orderTraversal(node.Left, s)
	*s = append(*s, node.Val)
	orderTraversal(node.Right, s)
}
