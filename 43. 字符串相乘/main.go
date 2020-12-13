package main

import (
	"fmt"
)

//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
//示例 1:
//
//输入: num1 = "2", num2 = "3"
//输出: "6"
//示例 2:
//
//输入: num1 = "123", num2 = "456"
//输出: "56088"
//说明：
//
//num1 和 num2 的长度小于110。
//num1 和 num2 只包含数字 0-9。
//num1 和 num2 均不以零开头，除非是数字 0 本身。
//不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/multiply-strings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(multiply("99", "99"))
}


// i + j
//   12
//   34
//    8
//   4
// 3 6
// 4 0 8

func multiply(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	data := make([]int, l1 + l2)
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			index := i + j + 1
			data[index] += int(num1[i] - '0') * int(num2[j] - '0')
		}
	}
	for i := l1 + l2 - 1; i >= 1; i-- {
		tmp := data[i]
		data[i] = tmp % 10
		data[i-1] += tmp / 10
	}
	res := []byte{}
	start := 0
	for i := 0; i < l1+l2; i++ {
		if data[i] != 0 {
			break
		}
		start = i + 1
	}
	for i := start; i < l1+l2; i++ {
		res = append(res, byte(data[i]) + '0')
	}
	if len(res) == 0 {
		return "0"
	}
	return string(res)
}
