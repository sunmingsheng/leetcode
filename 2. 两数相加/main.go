package main

import "math"

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

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head1 := l1
	head2 := l2
	list1 := []*ListNode{}
	list2 := []*ListNode{}
	flag1 := true
	flag2 := true
	for {
		if !flag1 && !flag2 {
			break
		}
		if head1 != nil {
			list1 = append(list1, head1)
			head1 = head1.Next
		} else {
			flag1 = false
		}
		if head2 != nil {
			list2 = append(list2, head2)
			head2 = head2.Next
		} else {
			flag2 = false
		}
	}
	len1 := len(list1)
	len2 := len(list2)
	if len1 != len2 {
		for len1 < len2 {
			list1 = append(list1, &ListNode{Val:0})
			len1 += 1
		}
		for len1 > len2 {
			list2 = append(list2, &ListNode{Val:0})
			len2 += 1
		}
	}
	data := 0
	for key, node := range list1 {
		data += node.Val * int(math.Pow10(len(list1) - key - 1))
	}
	for key, node := range list2 {
		data += node.Val * int(math.Pow10(len(list2) - key - 1))
	}

	results := []int{}
	i := 1
	for {
		res := data % 10
		results = append(results, res)
		data -= res
		data /= 10
		if data < 0 {
			break
		}
        i ++
	}

	root := &ListNode{
		Val:  results[0],
		Next: nil,
	}
	head := root
	for i := 1; i < len(results); i++ {
		head.Next = &ListNode{
			Val:  results[i],
			Next: nil,
		}
		head = head.Next
	}
	return root
}
