package main

//从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
//
// 
//
//例如:
//给定二叉树: [3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回：
//
//[3,9,20,15,7]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Queue []*TreeNode

func (q *Queue) Put(value *TreeNode) {
	*q = append(*q, value)
}

func (q *Queue) Pop() *TreeNode {
	if q.isEmpty() {
		return nil
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q *Queue) isEmpty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

func main() {
	left := &TreeNode{Val:1}
	right := &TreeNode{Val:3}
	head := &TreeNode{Val:2}
	head.Left = left
	head.Right = right
	levelOrder(head)
}

func levelOrder(root *TreeNode) []int {
	s := []int{}
	if root == nil {
		return s
	}
	q := &Queue{}
	q.Put(root)
	for !q.isEmpty() {
		node := q.Pop()
		s = append(s, node.Val)
		if node.Left != nil {
			q.Put(node.Left)
		}
		if node.Right != nil {
			q.Put(node.Right)
		}
	}
	return s
}
