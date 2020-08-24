package main

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//
//示例：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {

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
