package main

import "fmt"

//请实现有重复数字的有序数组的二分查找。
//输出在数组中第一个大于等于查找值的位置，如果数组中不存在这样的数，则输出数组长度加一。
//示例1
//输入
//复制
//5,4,[1,2,4,4,5]

func main() {
	fmt.Println(upper_bound_(5,4,[]int{1,2,4,4,5}))
}

/**
 * 二分查找
 * @param n int整型 数组长度
 * @param v int整型 查找值
 * @param a int整型一维数组 有序数组
 * @return int整型
 */
func upper_bound_( n int ,  v int ,  a []int ) int {
	if n == 0 || a[n-1] < v {
		return n + 1
	}
	l := 0
	r := n-1
	mid := l + (r-l)/2
	for l < r {
		if a[mid] >= v {
			r = mid
		} else {
			l = mid+1
		}
		mid = l + (r-l)/2
	}
	return mid+1
}

