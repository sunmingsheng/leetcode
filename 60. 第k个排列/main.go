package main

import (
	"fmt"
	"strconv"
)

//给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
//
//按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
//
//"123"
//"132"
//"213"
//"231"
//"312"
//"321"
//给定 n 和 k，返回第 k 个排列。
//
//说明：
//
//给定 n 的范围是 [1, 9]。
//给定 k 的范围是[1,  n!]。
//示例 1:
//
//输入: n = 3, k = 3
//输出: "213"
//示例 2:
//
//输入: n = 4, k = 9
//输出: "2314"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/permutation-sequence
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(getPermutation(4,9))
}

func getPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	}
	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, i+1)
	}
	factorialArr := factorial(n - 1)
	var s string
	var fun func(n, k int)

	fun = func(n, k int) {
		f := factorialArr[n-1]
		sh := k / f
		yu := k % f
		var index int
		if yu == 0 {
			index = sh - 1
			yu = k - index*f
		} else {
			index = sh
		}
		s += strconv.Itoa(arr[index])
		arr = append(arr[:index], arr[index+1:]...)
		if len(arr) == 1 {
			s += strconv.Itoa(arr[0])
			return
		}
		fun(n-1, yu)
	}

	fun(n, k)
	return s
}

func factorial(n int) []int {
	factorialArr := []int{
		1,
	}
	facVal := 1
	for i := 1; i <= n; i++ {
		facVal *= i
		factorialArr = append(factorialArr, facVal)
	}
	return factorialArr
}

