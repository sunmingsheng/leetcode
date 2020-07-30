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
	length int
	minValue int
	minIndex int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:     []int{},
		length:   0,
		minValue: 0,
		minIndex: -1,
	}
}


func (this *MinStack) Push(x int)  {
	this.data = append(this.data, x)
	this.length += 1
	if x < this.minValue {
		this.minValue =
	}
}


func (this *MinStack) Pop()  {

}


func (this *MinStack) Top() int {

}


func (this *MinStack) GetMin() int {

}

func main() {
	obj := Constructor();
	obj.Push(1)
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
