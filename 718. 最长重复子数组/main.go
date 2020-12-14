package main

import "fmt"

//给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。
//
// 
//
//示例：
//
//输入：
//A: [1,2,3,2,1]
//B: [3,2,1,4,7]
//输出：3
//解释：
//长度最长的公共子数组是 [3, 2, 1] 。
// 
//
//提示：
//
//1 <= len(A), len(B) <= 1000
//0 <= A[i], B[i] < 100
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(findLength([]int{1,2,3,2,1}, []int{3,2,1,4,7}))
}

func findLength(A []int, B []int) int {
	length := len(A)
	if length <= 0 {
		return 0
	}
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
	}
	max := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if A[i] != B[j] {
				dp[i][j] = 0
			} else {
				dp[i][j] = 1
				if i >= 1 && j >= 1 {
					dp[i][j] += dp[i-1][j-1]
				}
				if max < dp[i][j] {
					max = dp[i][j]
				}
			}
		}
	}
	return  max
}
