package main

import (
	"container/heap"
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

type MedianFinder struct {
	minHeap *heapMin
	maxHeap *heapMin
}

type heapMin []int

func (h heapMin) Len() int {
	return len(h)
}

func (h heapMin) Less(i,j int) bool {
	return h[i] < h[j]
}

func (h *heapMin) Push(x interface{}) {
	*h = append(*h,x.(int))
}

func (h *heapMin) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}

func (h *heapMin) Swap(i,j int) {
	(*h)[i],(*h)[j] = (*h)[j],(*h)[i]
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: new(heapMin),
		maxHeap: new(heapMin),
	}
}

func (this *MedianFinder) AddNum(num int)  {
	if this.minHeap.Len() > 0 && num > (*this.minHeap)[0] {
		heap.Push(this.minHeap,num)
	} else {
		heap.Push(this.maxHeap,-num)
	}

	if this.minHeap.Len() - this.maxHeap.Len() == 2 {
		heap.Push(this.maxHeap,-(heap.Pop(this.minHeap)).(int))
	} else if this.maxHeap.Len() - this.minHeap.Len() == 2 {
		heap.Push(this.minHeap,-(heap.Pop(this.maxHeap)).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() > this.maxHeap.Len() {
		return float64((*this.minHeap)[0])
	} else if this.minHeap.Len() < this.maxHeap.Len() {
		return -float64((*this.maxHeap)[0])
	}
	return float64((*this.minHeap)[0]-(*this.maxHeap)[0])/float64(2)
}




func main() {


	//obj := Constructor()
	//obj.AddNum(1)
	//fmt.Println(obj.FindMedian())
}


