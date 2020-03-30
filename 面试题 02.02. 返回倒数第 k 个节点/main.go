package main

import "fmt"

//实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
//
//注意：本题相对原题稍作改动
//
//示例：
//
//输入： 1->2->3->4->5 和 k = 2
//输出： 4
//说明：
//
//给定的 k 保证是有效的。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	head := &ListNode{Val:1}
	fmt.Println(kthToLast(head, 1))
}


func kthToLast(head *ListNode, k int) int {
	s := []int{}
	node := head
	for node != nil {
		s = append(s, node.Val)
		node = node.Next
	}
	return s[len(s) - k]
}