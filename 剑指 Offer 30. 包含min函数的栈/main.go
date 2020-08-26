package main

import "fmt"

//定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
//
// 
//
//示例:
//
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.min();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.min();   --> 返回 -2.
// 
//
//提示：
//
//各函数的调用总次数不超过 20000 次
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type MinStack struct {
	length int
	data []int
	minData []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		length:  0,
		data:    []int{},
		minData: []int{},
	}
}


func (this *MinStack) Push(x int)  {
	this.data = append(this.data, x)
	if this.length == 0 {
		this.minData = append(this.minData, x)
	} else {
		if x > this.minData[this.length - 1] {
			this.minData = append(this.minData, this.minData[this.length - 1])
		} else {
			this.minData = append(this.minData, x)
		}
	}
	this.length += 1
}


func (this *MinStack) Pop()  {
	if this.length <= 0 {
		return
	}
	this.data = this.data[:this.length-1]
	this.minData = this.minData[:this.length-1]
	this.length -= 1
}


func (this *MinStack) Top() int {
	if this.length <= 0 {
		return 0
	}
	return this.data[this.length-1]
}


func (this *MinStack) Min() int {
	if this.length <= 0 {
		return 0
	}
	return this.minData[this.length-1]
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
	fmt.Println(obj.Min())
}