package main

import (
	"fmt"
	"sort"
)

//给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。
//
// 
//
//示例：
//
//输入：[1,2,3,4,5,null,7,8]
//
//1
///  \
//2    3
/// \    \
//4   5    7
///
//8
//
//输出：[[1],[2,3],[4,5,7],[8]]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/list-of-depth-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
    Left *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val int
    Next *ListNode
}

func main() {
	head := TreeNode{Val:1}
	right := TreeNode{Val:2}
	left := TreeNode{Val:3}
	head.Left = &left
	head.Right = &right
	results := listOfDepth(&head)
	for _, head := range results {
		fmt.Println("------------------------")
		current := head
		for current != nil {
			fmt.Println(current.Val)
			current = current.Next
		}
	}
}

func listOfDepth(tree *TreeNode) []*ListNode {
	resultMap := make(map[int][]*TreeNode)
	resultMap[0] = []*TreeNode{tree}
	resultLists := []*ListNode{}
	printLevelTree(tree, 1, resultMap)
	//整理数据
	mapSort := []int{}
	for key := range resultMap {
		mapSort = append(mapSort, key)
	}
	sort.Ints(mapSort)
	for _, mapKey := range mapSort {
		treeNodes := resultMap[mapKey]
		var beforeListNode *ListNode
		var headListNode *ListNode
		if len(treeNodes) > 0 {
			for _, treeNode := range treeNodes {
				currentListNode := &ListNode{Val: treeNode.Val}
				if beforeListNode != nil {
					beforeListNode.Next = currentListNode
				} else {
					headListNode = currentListNode
				}
				beforeListNode = currentListNode
			}
		}
		resultLists = append(resultLists, headListNode)
	}
	return resultLists
}

func printLevelTree(node *TreeNode, deep int, results map[int][]*TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil || node.Right != nil {
		result, ok := results[deep]
		if !ok {
			result = []*TreeNode{}
		}
		if node.Left != nil {
			result = append(result, node.Left)
		}
		if node.Right != nil {
			result = append(result, node.Right)
		}
		results[deep] = result
		printLevelTree(node.Left, deep+1, results)
		printLevelTree(node.Right, deep+1, results)
	}
}
