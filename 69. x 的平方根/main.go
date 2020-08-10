package main

import "fmt"

//实现 int sqrt(int x) 函数。
//
//计算并返回 x 的平方根，其中 x 是非负整数。
//
//由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
//
//示例 1:
//
//输入: 4
//输出: 2
//示例 2:
//
//输入: 8
//输出: 2
//说明: 8 的平方根是 2.82842...,
//     由于返回类型是整数，小数部分将被舍去。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sqrtx
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println("结果:", mySqrt(4))
}

//func mySqrt(x int) int {
//	if x <= 1 {
//		return x
//	}
//	for i := 1; i <= x; i++ {
//		res := i * i
//		if res == x {
//			return i
//		}
//		if res > x {
//			return i - 1
//		}
//	}
//	return 0
//}

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	l := 0
	r := x
	middle := 0
	ans := 0
	for l <= r {
		middle = l + (r - l) / 2
		res := middle * middle
		if res == x {
			return middle
		}
		if res > x {
			r = middle - 1
			ans = middle - 1
		} else {
			l = middle + 1
		}
	}
	return ans
}
