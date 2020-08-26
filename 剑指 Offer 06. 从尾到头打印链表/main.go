package main

//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
//示例 1：
//
//输入：head = [1,3,2]
//输出：[2,3,1]

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {

}

func reversePrint(head *ListNode) []int {
	s := []int{}
	traversal(head, &s)
	res := make([]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		res[len(s) - i - 1] = s[i]
	}
	return res
}

func traversal(node *ListNode, s *[]int) {
	for node != nil {
		*s = append(*s, node.Val)
		node = node.Next
	}
}