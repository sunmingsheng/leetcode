package main

import "fmt"

//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//
//示例 1:
//
//输入: 1->1->2
//输出: 1->2
//示例 2:
//
//输入: 1->1->2->3->3
//输出: 1->2->3
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	fmt.Println(deleteDuplicates(&ListNode{Val:1}))
}

func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for {
		if node == nil || node.Next == nil {
			break
		}
		if node.Val == node.Next.Val {
			temp := node.Next
			node.Next = temp.Next
			temp.Next = nil
		} else {
			node = node.Next
		}
	}
	return head
}
