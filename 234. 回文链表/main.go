package main

import "fmt"

//请判断一个链表是否为回文链表。
//
//示例 1:
//
//输入: 1->2
//输出: false
//示例 2:
//
//输入: 1->2->2->1
//输出: true
//进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/palindrome-linked-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	fmt.Println(isPalindrome(&ListNode{Val:1}))
}

func isPalindrome(head *ListNode) bool {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	length := len(arr)
	if length <= 1 {
		return true
	}
	for i := 0; i < length/2;i ++ {
		j := length - i - 1
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}
