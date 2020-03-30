package main

import "fmt"

//满二叉树是一类二叉树，其中每个结点恰好有 0 或 2 个子结点。
//
//返回包含 N 个结点的所有可能满二叉树的列表。 答案的每个元素都是一个可能树的根结点。
//
//答案中每个树的每个结点都必须有 node.val=0。
//
//你可以按任何顺序返回树的最终列表。
//
//示例：
//
//输入：7
//输出：[[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/all-possible-full-binary-trees
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(allPossibleFBT(7))
}

func allPossibleFBT(N int) []*TreeNode {
	if N == 0 || N % 2 == 0 {
		return []*TreeNode{}
	}
	results := []*TreeNode{}
	root := &TreeNode{Val:0}
	results = append(results, root)
	return results
}
