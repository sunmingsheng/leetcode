package main

//给定一个二叉树，确定它是否是一个完全二叉树。
//
//百度百科中对完全二叉树的定义如下：
//
//若设二叉树的深度为 h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全二叉树。（注：第 h 层可能包含 1~ 2h 个节点。）
//
// 
//
//示例 1：
//
//
//
//输入：[1,2,3,4,5,6]
//输出：true
//解释：最后一层前的每一层都是满的（即，结点值为 {1} 和 {2,3} 的两层），且最后一层中的所有结点（{4,5,6}）都尽可能地向左。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	nodes := make(map[int][]*TreeNode)
	levelTravel(root, 1, nodes)
	maxLevel := 1
	for key, _ := range nodes {
		if key > maxLevel {
			maxLevel = key
		}
	}

	for key, lNodes := range nodes {
		if key < maxLevel - 1 {
			for _, node := range lNodes {
				if node == nil {
					return false
				}
			}
		} else if key == maxLevel - 1 {
			flag := true
			for _, node := range lNodes {
				if node == nil && flag {
					flag = false
				}
				if node != nil && !flag {
					return false
				}
			}
		}
	}
	return true
}

func levelTravel(node *TreeNode, level int, nodes map[int][]*TreeNode) {
	nodes[level] = append(nodes[level], node)
	if node != nil {
		levelTravel(node.Left, level+1, nodes)
		levelTravel(node.Right, level+1, nodes)
	}
}
