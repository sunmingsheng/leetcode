package main

//判断给定的链表中是否有环
//扩展：
//你能给出空间复杂度的解法么？

type ListNode struct{
	Val int
	Next *ListNode
}


/**
 *
 * @param head ListNode类
 * @return bool布尔型
 */
func hasCycle( head *ListNode ) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
