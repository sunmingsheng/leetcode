package main

//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
//
// 
//
//示例 1：
//
//输入：
//["CQueue","appendTail","deleteHead","deleteHead"]
//[[],[3],[],[]]
//输出：[null,null,3,-1]
//示例 2：
//
//输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type CQueue struct {
	length int
	stack_1 []int
	stack_2 []int
}


func Constructor() CQueue {
	return CQueue{
		stack_1: []int{},
		stack_2: []int{},
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.length += 1
	this.stack_1 = append(this.stack_1, value)
}


func (this *CQueue) DeleteHead() int {
	if this.length <= 0 {
		return -1
	}
	length := len(this.stack_2)
	if length == 0 {
		for i := len(this.stack_1) - 1; i >= 0; i-- {
			this.stack_2 = append(this.stack_2, this.stack_1[i])
		}
		this.stack_1 = []int{}
	}
	this.length -= 1
	temp := make([]int, len(this.stack_2) - 1)
	copy(temp, this.stack_2[0:len(this.stack_2) - 1])
	value := this.stack_2[len(this.stack_2) - 1]
	this.stack_2 = temp
	return value
}

func main() {
	obj := Constructor()
	obj.AppendTail(1)
	obj.DeleteHead()
}
