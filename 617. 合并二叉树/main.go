package main

import "fmt"

//给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
//
//你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。
//
//示例 1:
//
//输入:
//Tree 1                     Tree 2
//1                         2
/// \                       / \
//3   2                     1   3
///                           \   \
//5                             4   7
//输出:
//合并后的树:
//3
/// \
//4   5
/// \   \
//5   4   7
//注意: 合并必须从两个树的根节点开始。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-two-binary-trees
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	fmt.Println(mergeTrees(node1, node2))
}

type Queue []*TreeNode

func (q *Queue) Put(value *TreeNode) {
	*q = append(*q, value)
}

func (q *Queue) Pop() *TreeNode {
	result := (*q)[0]
	*q = (*q)[1:]
	return result
}

func (q *Queue) IsEmpty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	//通过层序遍历将树扫描到切片中
	s1 := &[]int{}
	levelTravel(t1, s1)
	s2 := &[]int{}
	levelTravel(t2, s2)
	//两个切片数据相加得到新的切片
	s3 := make([]int, len(*s1))
	copy(s3, *s1)
	for key, value := range *s2 {
		if key >= len(s3) {
			s3 = append(s3, value)
		} else {
			s3[key] += value
		}
	}
	//然后将新的切片渲染成树
	newRoot := &TreeNode{Val:s3[0]}
	ParseSliceToTree(newRoot, 0, s3)
	return newRoot
}

func levelTravel(head *TreeNode, s *[]int) {
	q := &Queue{}
	q.Put(head)
	for !q.IsEmpty() {
		node := q.Pop()
		if node != nil {
			*s = append(*s, node.Val)
		} else {
			*s = append(*s, 0)
		}
		if node != nil {
			//if node.Left != nil {
			q.Put(node.Left)
			//}
			//if node.Right != nil {
			q.Put(node.Right)
			//}
		}
	}
}

func ParseSliceToTree(node *TreeNode, index int, s []int) {
	if node == nil {
		return
	}
	//寻找子节点
	leftIndex := 2 * index + 1
	if leftIndex >= len(s) {
		return
	} else {
		node.Left = &TreeNode{Val:s[leftIndex]}
		ParseSliceToTree(node.Left, leftIndex, s)
	}
	rightIndex := 2 * index + 2
	if rightIndex >= len(s) {
		return
	} else {
		node.Right = &TreeNode{Val:s[rightIndex]}
		ParseSliceToTree(node.Right, rightIndex, s)
	}
}
