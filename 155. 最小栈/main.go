package main

import "fmt"

//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) —— 将元素 x 推入栈中。
//pop() —— 删除栈顶的元素。
//top() —— 获取栈顶元素。
//getMin() —— 检索栈中的最小元素。
// 
//
//示例:
//
//输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/min-stack
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type MinStack struct {
	data []int
	minData []int
	length int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:   []int{},
		minData: []int{},
		length:   0,
	}
}


func (this *MinStack) Push(x int)  {
	this.data = append(this.data, x)
	if this.length == 0 {
		this.minData = append(this.minData, x)
	} else {
		if x < this.minData[this.length - 1] {
			this.minData = append(this.minData, x)
		} else {
			this.minData = append(this.minData, this.minData[this.length - 1])
		}
	}
	this.length += 1
}


func (this *MinStack) Pop()  {
	if this.length > 0 {
		this.data = this.data[0:this.length-1]
		this.minData = this.minData[0:this.length-1]
		this.length -= 1
	}
}


func (this *MinStack) Top() int {
	if this.length > 0 {
		return this.data[this.length-1]
	}
	return -1
}


func (this *MinStack) GetMin() int {
	if this.length > 0 {
		return this.minData[this.length-1]
	}
	return -1
}

func main() {
	obj := Constructor()
	obj.Push(0)
	obj.Push(3)
	obj.Push(4)
	obj.Push(1)
	obj.Push(2)
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
