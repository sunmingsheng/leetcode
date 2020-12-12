package main

import "fmt"

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-two-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
    l1 := &ListNode{Val:2, Next:&ListNode{Val:4}}
	l2 := &ListNode{Val:5, Next:&ListNode{Val:6, Next:&ListNode{Val:4}}}
	res := addTwoNumbers(l1, l2)
	printList(res)
}

func printList(node *ListNode) {
	for {
		if node == nil {
			break
		}
		fmt.Println(node.Val)
		node = node.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	prev := &ListNode{}
	s := prev
	nextVal := 0
	for {
		val := 0
		if l1 == nil && l2 == nil {
			break
		} else if l1 == nil || l2 == nil {
			if l1 != nil {
				val = l1.Val
			} 
			if l2 != nil {
				val = l2.Val
			}
			val += nextVal
			if val >= 10 {
				nextVal = 1
				val = val - 10
			} else {
				nextVal = 0
			}
			s.Next = &ListNode{Val: val}
			s = s.Next
			if l1 != nil {
				l1 = l1.Next
			}
			if l2 != nil {
				l2 = l2.Next
			}
		} else {
			val = l1.Val + l2.Val + nextVal
			if val >= 10 {
				nextVal = 1
				val = val - 10
			} else {
				nextVal = 0
			}
			s.Next = &ListNode{Val: val}
			s = s.Next
			l1 = l1.Next
			l2 = l2.Next
		}
	}
	return prev.Next
}
