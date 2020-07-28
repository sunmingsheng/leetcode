package main

type ListNode struct {
	Val int
	Next *ListNode
}


func main() {
	root := &ListNode{Val:1}
	deleteNode(root)
}

func deleteNode(node *ListNode) {
	if node == nil {
		return
	}
	fastNode := node
	slowNode := node
	prevNode := node
	for fastNode != nil && fastNode.Next != nil {
		fastNode = fastNode.Next.Next
		prevNode = slowNode
		slowNode = slowNode.Next
	}
	prevNode.Next = slowNode.Next
}
