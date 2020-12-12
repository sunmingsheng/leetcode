package main

import (
	"fmt"
	"sort"
)

//栈排序。 编写程序，对栈进行排序使最小元素位于栈顶。最多只能使用一个其他的临时栈存放数据，但不得将元素复制到别的数据结构（如数组）中。该栈支持如下操作：push、pop、peek 和 isEmpty。当栈为空时，peek 返回 -1。
//
//示例1:
//
//输入：
//["SortedStack", "push", "push", "peek", "pop", "peek"]
//[[], [1], [2], [], [], []]
//输出：
//[null,null,null,1,null,2]
//示例2:
//
//输入：
//["SortedStack", "pop", "pop", "push", "pop", "isEmpty"]
//[[], [], [], [1], [], []]
//输出：
//[null,null,null,null,null,true]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sort-of-stacks-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type SortedStack struct {
	length int
	data []int
}

func Constructor() SortedStack {
	return SortedStack{
		length: 0,
		data:   []int{},
	}
}


func (this *SortedStack) Push(val int)  {
	this.length += 1
	this.data = append(this.data, val)
	sort.Ints(this.data)
}


func (this *SortedStack) Pop()  {
	if this.length > 0 {
		this.data = this.data[1:]
		this.length -= 1
	}
}


func (this *SortedStack) Peek() int {
	if this.length <= 0 {
		return -1
	}
	return this.data[0]
}


func (this *SortedStack) IsEmpty() bool {
	if this.length <= 0 {
		return true
	}
	return false
}

func main() {
	obj := Constructor();
	obj.Push(1212)
	obj.Pop()
	fmt.Println(obj.Peek())
	fmt.Println(obj.IsEmpty())
}

