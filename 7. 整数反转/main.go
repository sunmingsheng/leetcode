package main

import (
	"fmt"
	"math"
)

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
//示例 1:
//
//输入: 123
//输出: 321
// 示例 2:
//
//输入: -123
//输出: -321
//示例 3:
//
//输入: 120
//输出: 21
//注意:
//
//假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reverse-integer
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(reverse(-120))
}

func reverse(x int) int {
	if x == 0 {
		return x
	}
    flag := false
    if x < 0 {
    	flag = true
    	x = 0 - x
	}
	c := 10
	data := []int{}
	for {
		if x <= 0 {
			break
		}
		tmp := x % c
		data = append(data, tmp * 10 / c)
		x = x - tmp
		c = c * 10
	}
	val := 0
	for i := 0; i < len(data); i++ {
		val = val * 10 + data[i]
	}
	if flag {
		val = 0 - val
		if val < math.MinInt32 {
			return 0
		}
	} else {
		if val > math.MaxInt32 {
			return 0
		}
	}
	return val
}