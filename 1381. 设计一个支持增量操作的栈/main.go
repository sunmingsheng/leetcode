package main

import "fmt"

func main() {
	c := Constructor(1)
	fmt.Println(c)
}

type CustomStack struct {
	maxSize int
	size int
	s *[]int
}

func Constructor(maxSize int) CustomStack {
	s := &[]int{}
	return CustomStack{
		maxSize: maxSize,
		size:    0,
		s:       s,
	}
}


func (this *CustomStack) Push(x int)  {
	if this.maxSize > this.size {
		*this.s = append(*this.s, x)
	}
}


func (this *CustomStack) Pop() int {
	if this.size > 0 {
		r := (*this.s)[0]
		this.size -= 1
		(*this.s) = (*this.s)[1:]
		return r
	}
	return -1
}


func (this *CustomStack) Increment(k int, val int)  {
	if k > this.size {
		k = this.size
	}
	for i := 0; i < k; i++ {
		(*this.s)[i] += val
	}
}
