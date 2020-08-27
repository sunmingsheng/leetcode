package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
//
//例如，
//
//[2,3,4] 的中位数是 3
//
//[2,3] 的中位数是 (2 + 3) / 2 = 2.5
//
//设计一个支持以下两种操作的数据结构：
//
//void addNum(int num) - 从数据流中添加一个整数到数据结构中。
//double findMedian() - 返回目前所有元素的中位数。
//示例：
//
//addNum(1)
//addNum(2)
//findMedian() -> 1.5
//addNum(3)
//findMedian() -> 2
//进阶:
//
//如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
//如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
//通过次数21,095提交次数44,212
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-median-from-data-stream
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


type PriorityQueue []int

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i] > (*pq)[j]
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Pop() interface{} {
	temp := (*pq)[len(*pq) - 1]
	*pq = (*pq)[:len(*pq)-1]
	return temp
}

func (pq *PriorityQueue) Push(item interface{})  {
	*pq = append(*pq, item.(int))
}

type MedianFinder struct {
	length int
	data []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		length: 0,
		data:   []int{},
	}
}

func (this *MedianFinder) AddNum(num int)  {
	this.data = append(this.data, num)
	this.length += 1
}

func (this *MedianFinder) FindMedian() float64 {
	sort.Ints(this.data)
	if this.length % 2 == 1 {
		return float64(this.data[this.length/2])
	} else {
		return float64(this.data[this.length/2] + this.data[this.length/2 - 1]) / 2
	}
}


func main() {
	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq,1)
	heap.Push(pq,3)
	heap.Push(pq,2)

	fmt.Println(heap.Pop(pq))
	fmt.Println(heap.Pop(pq))
	fmt.Println(heap.Pop(pq))


	//obj := Constructor()
	//obj.AddNum(1)
	//fmt.Println(obj.FindMedian())
}


