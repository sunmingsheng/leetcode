package main

import "fmt"

//给定一个未经排序的整数数组，找到最长且连续的的递增序列。
//
//示例 1:
//
//输入: [1,3,5,         n 4,7]
//输出: 3
//解释: 最长连续递增序列是 [1,3,5], 长度为3。
//尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为5和7在原数组里被4隔开。
//示例 2:
//
//输入: [2,2,2,2,2]
//输出: 1
//解释: 最长连续递增序列是 [2], 长度为1。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(findLengthOfLCIS([]int{2,2,2,2}))
}

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxLen := 0
	res := 0
	for i := 0; i < len(nums) - 1; i++ {
		res += 1
		if nums[i] >= nums[i + 1] {
			if res > maxLen {
				maxLen = res
			}
			res = 0
		}
	}
	if res + 1 > maxLen {
		maxLen = res + 1
	}
	return maxLen
}