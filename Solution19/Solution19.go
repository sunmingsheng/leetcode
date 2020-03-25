package main

import "fmt"

type ListNode19 struct {
	Val int
	Next *ListNode19
}

func main() {
	head := &ListNode19{Val:1, Next:nil}
	fmt.Println(removeNthFromEnd(head, 1))
}


func removeNthFromEnd(head *ListNode19, n int) *ListNode19 {
	oriHead := head
	s := []*ListNode19{}
	for head != nil {
		s = append(s, head)
		head = head.Next
	}
	if len(s) <= n {
		return oriHead.Next
	} else {
		s[len(s) - n - 1].Next = s[len(s) - n].Next
		s[len(s) - n].Next = nil
	}
	return oriHead
}
