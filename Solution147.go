package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	p := new(ListNode)
	p.Val = 1
	insertionSortList(p)
}

func insertionSortList(head *ListNode) *ListNode {
	currentNode := head
	for currentNode != nil {
		nextNode := currentNode.Next
		if nextNode == nil {
			break
		} else {
			//将nextnode与其之前的所有有序的节点进行比较排序
			tempNode := head
			for tempNode != nextNode {
				if tempNode.Val <= nextNode.Val {
					tempNode = tempNode.Next
				} else {
                    //还需要记录befor和next节点信息
				}
			}
			currentNode = nextNode
		}
	}
	return head
}