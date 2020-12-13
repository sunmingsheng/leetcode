package main

//给定两个用链表表示的整数，每个节点包含一个数位。
//
//这些数位是反向存放的，也就是个位排在链表首部。
//
//编写函数对这两个整数求和，并用链表形式返回结果。
//
// 
//
//示例：
//
//输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
//输出：2 -> 1 -> 9，即912
//进阶：思考一下，假设这些数位是正向存放的，又该如何解决呢?
//
//示例：
//
//输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
//输出：9 -> 1 -> 2，即912
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-lists-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

}

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//链表相加
	dummyHead := &ListNode{}
	currHead := dummyHead
	for {
		if l1 == nil || l2 == nil {
			break
		}
		val := l1.Val + l2.Val
		currHead.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		currHead = currHead.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil {
		currHead.Next = l2
	}
	if l2 == nil {
		currHead.Next = l1
	}
	//进制处理
	currHead = dummyHead
	for {
		if currHead == nil {
			break
		}
		if currHead.Val >= 10 {
			nextVal := 1
			currHead.Val -= 10
			if currHead.Next != nil {
				currHead.Next.Val += nextVal
			} else {
				currHead.Next = &ListNode{
					Val:  nextVal,
					Next: nil,
				}
			}
		}
		currHead = currHead.Next
	}
	//再翻转输出
	return dummyHead.Next
}

//翻转
func reset(l *ListNode) *ListNode {
	//7 -> 2 -> 4 -> 3
	//3 -> 4 -> 2 -> 7
	if l == nil {
		return l
	}
	var prev *ListNode
	curr := l
	for {
		if curr == nil {
			break
		}
		next := curr.Next //p:0 c:7 n:2
		curr.Next = prev  //
		prev = curr
		curr = next
	}
	return prev
}
