package main

import "fmt"

//
//输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//
// 
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
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(buildTree([]int{},[]int{}))
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