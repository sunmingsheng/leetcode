package main

import "fmt"

//给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
//例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回锯齿形层次遍历如下：
//
//[
//[3],
//[20,9],
//[15,7]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:1, Left:&TreeNode{Val:-1}, Right:&TreeNode{Val:4}}
	fmt.Println(zigzagLevelOrder(root))
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := make(map[int][]*TreeNode)
	q[0] = append(q[0], root)
	print(root.Left, 1, &q)
	print(root.Right, 1, &q)
	maxValue := 0
	for key, _ := range q {
		if maxValue < key {
			maxValue = key
		}
	}
	//遍历循环q，如果索引为奇数，则翻转
	results := make([][]int, len(q))
	for i := 0; i <= maxValue; i++ {
		key := i
		items := q[key]
		if key % 2 == 0 {
			for _, node := range items {
				results[key] = append(results[key], node.Val)
			}
		} else {
			for i := len(items) - 1; i >=0 ; i-- {
				results[key] = append(results[key], items[i].Val)
			}
		}
	}
	return results
}

func print(node *TreeNode, level int, q *map[int][]*TreeNode) {
	if node == nil {
		return
	}
	(*q)[level] = append((*q)[level], node)
	print(node.Left, level + 1, q)
	print(node.Right, level + 1, q)
}


