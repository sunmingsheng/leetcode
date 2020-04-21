package main

import "fmt"

//给定一个未经排序的整数数组，找到最长的递增序列。
//
//示例 1:
//
//输入: [1,3,5,4,7]
//输出: 3
//解释: 最长连续递增序列是 [1,3,5,7], 长度为4。

func main() {
	fmt.Println(findLengthOfLCIS([]int{1,3,5,4,7}))
}

// 1,3,5,4,7
func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	for key, _ := range dp {
		dp[key]	= 1
	}
	for  i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j] + 1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	res := 0
	for _, value := range dp {
		if value > res {
			res = value
		}
	}
	return res
}
