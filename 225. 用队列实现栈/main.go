package main

import "fmt"

//使用队列实现栈的下列操作：
//
//push(x) -- 元素 x 入栈
//pop() -- 移除栈顶元素
//top() -- 获取栈顶元素
//empty() -- 返回栈是否为空
//注意:
//
//你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
//你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
//你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/implement-stack-using-queues
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type MyStack struct {
	size int
	list1 []int
	list2 []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		size:  0,
		list1: []int{},
		list2: []int{},
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.list1 = append(this.list1, x)
	this.size += 1
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.size < 1 {
		return -1
	}
	var data int
	data = this.list1[this.size - 1]
	this.list1 = this.list1[0:this.size - 1]
	this.size -= 1
	return data
}


/** Get the top element. */
func (this *MyStack) Top() int {
	if this.size < 1 {
		return -1
	}
	var data int
	data = this.list1[this.size - 1]
	return data
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.size <= 0 {
		return true
	}
	return false
}


func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Top())
	fmt.Println(obj.Empty())
}