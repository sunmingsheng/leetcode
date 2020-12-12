package main

//给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
//
//示例 1:
//
//输入: 1->2->3->4->5->NULL, k = 2
//输出: 4->5->1->2->3->NULL
//解释:
//向右旋转 1 步: 5->1->2->3->4->NULL
//向右旋转 2 步: 4->5->1->2->3->NULL
//示例 2:
//
//输入: 0->1->2->NULL, k = 4
//输出: 2->0->1->NULL
//解释:
//向右旋转 1 步: 2->0->1->NULL
//向右旋转 2 步: 1->2->0->NULL
//向右旋转 3 步: 0->1->2->NULL
//向右旋转 4 步: 2->0->1->NULL
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/rotate-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	length := 0
	curr := head
	var tail *ListNode
	for curr != nil {
		tail = curr
		length += 1
		curr = curr.Next
	}
	k = k % length
	if k == 0 {
		return head
	}
	var newHead *ListNode
	curr = head
	index := 0
	for curr != nil {
		index += 1
		if length - k == index {
			newHead = curr.Next
			tail.Next = head
			curr.Next = nil
			break
		}
		curr = curr.Next
	}
	return newHead
}