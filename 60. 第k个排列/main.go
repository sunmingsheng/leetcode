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
	res := []string{}
	strArr := make([]string, n)
	for i := 1; i <= n; i++ {
		strArr[i-1] = strconv.Itoa(i)
	}
	dfs(strArr, "", &res, k)
	return res[k-1]
}

func dfs(strArr []string, tmp string,  res *[]string, targetNum int) {
	if len(strArr) == 0 {
		*res = append(*res, tmp)
		return
	}
	if len(*res) > targetNum {
		return
	}
	for i := 0; i < len(strArr); i++ {
		data := make([]string, i)
		copy(data, strArr[:i])
		data = append(data, strArr[i+1:]...)
		dfs(data, tmp + strArr[i], res, targetNum)
	}
}
