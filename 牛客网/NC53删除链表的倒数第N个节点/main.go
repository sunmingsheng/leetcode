package main


type ListNode struct{
	Val int
	Next *ListNode
}

/**
 *
 * @param head ListNode类
 * @param n int整型
 * @return ListNode类
 */
func removeNthFromEnd( head *ListNode ,  n int ) *ListNode {
	//0->1->2->3->4 n=2
	dummy := &ListNode{Val:-1}
	dummy.Next = head
	fast := dummy
	slow := dummy
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
