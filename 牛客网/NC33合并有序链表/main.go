package main


type ListNode struct{
	Val int
   Next *ListNode
}

/**
 *
 * @param l1 ListNode类
 * @param l2 ListNode类
 * @return ListNode类
 */
func mergeTwoLists( l1 *ListNode ,  l2 *ListNode ) *ListNode {
	dummy := &ListNode{}
	current  := dummy
	for (l1 != nil && l2 != nil) {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}
	return dummy.Next
}