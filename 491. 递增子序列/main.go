package main

import "fmt"

//给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
//
//示例:
//
//输入: [4, 6, 7, 7]
//输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
//说明:
//
//给定数组的长度不会超过15。
//数组中的整数范围是 [-100,100]。
//给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/increasing-subsequences
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main()  {
	fmt.Println(findSubsequences([]int{1,3,5,7}))
}

func findSubsequences(nums []int) [][]int {
	res  := [][]int{}
	dfs(nums, 0, &res, []int{})
	return res
}

func dfs(nums []int, index int, res *[][]int, list []int) {
	if len(list) >= 2 {
		dest := make([]int, len(list))
		copy(dest, list)
		*res = append(*res, dest)
	}
	m := make(map[int]struct{})
	for i := index; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		if len(list) == 0 || nums[i] >= list[len(list)-1] {
			m[nums[i]] = struct{}{}
			dfs(nums, i+1, res, append(list, nums[i]))
		}
	}
}
