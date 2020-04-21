package main

import "fmt"

//给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
//
//本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
//
//示例:
//
//给定的有序链表： [-10, -3, 0, 5, 9],
//
//一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
//
//0
/// \
//-3   9
///   /
//-10  5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val int
	Next *ListNode
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	head := &ListNode{Val:1,Next:&ListNode{Val:2}}
	fmt.Println(sortedListToBST(head))
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{
					Val:head.Val,}
	}
	fastNode := head
	slowNode := head
	prevNode := head
	//-10, -3, 0, 5, 9, 10
	for fastNode != nil && fastNode.Next != nil {
		fastNode = fastNode.Next.Next
		prevNode = slowNode
		slowNode = slowNode.Next
	}
	middleNode := slowNode
	prevNode.Next = nil
	if prevNode != nil {
		prevNode.Next = nil
	}
    node := &TreeNode{
		Val:   middleNode.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(middleNode.Next),
	}
	return node
}
