package main

import "fmt"

//给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
//
//你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
// 
//
//进阶：
//
//如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
//
// 
//
//示例：
//
//输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 8 -> 0 -> 7
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-two-numbers-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val:9,Next:&ListNode{Val:2}}
	l2 := &ListNode{Val:1,Next:&ListNode{Val:2,Next:&ListNode{Val:3}}}
	print(addTwoNumbers(l1,l2))
}

func print(l *ListNode) {
	for {
		if l == nil {
			break
		}
		fmt.Println(l.Val)
		l = l.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//链表翻转
	l1 = reset(l1)
	l2 = reset(l2)
	//链表相加
	dummyHead := &ListNode{}
	currHead := dummyHead
	for {
		if l1 == nil || l2 == nil {
			break
		}
		val := l1.Val + l2.Val
		currHead.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		currHead = currHead.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil {
		currHead.Next = l2
	}
	if l2 == nil {
		currHead.Next = l1
	}
	//进制处理
	currHead = dummyHead
	for {
		if currHead == nil {
			break
		}
		if currHead.Val >= 10 {
			nextVal := 1
			currHead.Val -= 10
			if currHead.Next != nil {
				currHead.Next.Val += nextVal
			} else {
				currHead.Next = &ListNode{
					Val:  nextVal,
					Next: nil,
				}
			}
		}
		currHead = currHead.Next
	}
	//再翻转输出
	return reset(dummyHead.Next)
}

//翻转
func reset(l *ListNode) *ListNode {
	//7 -> 2 -> 4 -> 3
	//3 -> 4 -> 2 -> 7
	if l == nil {
		return l
	}
	var prev *ListNode
	curr := l
	for {
		if curr == nil {
			break
		}
		next := curr.Next //p:0 c:7 n:2
		curr.Next = prev  //
		prev = curr
		curr = next
	}
	return prev
}