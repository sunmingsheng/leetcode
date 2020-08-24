package main

import "math"

//给你一个链表数组，每个链表都已经按升序排列。
//
//请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
// 
//
//示例 1：
//
//输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//1->4->5,
//1->3->4,
//2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//示例 2：
//
//输入：lists = []
//输出：[]
//示例 3：
//
//输入：lists = [[]]
//输出：[]
// 
//
//提示：
//
//k == lists.length
//0 <= k <= 10^4
//0 <= lists[i].length <= 500
//-10^4 <= lists[i][j] <= 10^4
//lists[i] 按 升序 排列
//lists[i].length 的总和不超过 10^4
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {

}

//最小堆的解法更好
func mergeKLists(lists []*ListNode) *ListNode {
	res := &ListNode{Val:math.MinInt64}
	for _, list := range lists {
		res = mergeTwoLists(res, list)
	}
	return res.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	currHead := dummyHead
	for l1 != nil && l2 != nil {
		node := &ListNode{
			Val:  0,
			Next: nil,
		}
		if l1.Val < l2.Val {
			node.Val = l1.Val
			l1 = l1.Next
		} else {
			node.Val = l2.Val
			l2 = l2.Next
		}
		currHead.Next = node
		currHead = currHead.Next
	}
	if l1 != nil {
		currHead.Next = l1
	}
	if l2 != nil {
		currHead.Next = l2
	}
	return dummyHead.Next
}
