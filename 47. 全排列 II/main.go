package main

import (
	"fmt"
	"sort"
)

//给定一个可包含重复数字的序列，返回所有不重复的全排列。
//
//示例:
//
//输入: [1,1,2]
//输出:
//[
//[1,1,2],
//[1,2,1],
//[2,1,1]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/permutations-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	// 1 1 1 2 2
	fmt.Println(permuteUnique([]int{1,1,1,2}))
}

func permuteUnique(nums []int) [][]int {
	results := [][]int{}
	result := []int{}
	sort.Ints(nums)
	dfs(nums, &results, result)
	return results
}

func dfs(nums []int, results *[][]int, result []int) {
	if len(nums) <= 0 {
		*results = append(*results, result)
		return
	}
	for key, value := range nums {
		if key > 0 && nums[key - 1] == value {
			continue
		}
		data := append([]int{}, nums[0:key]...)
		data = append(data, nums[key + 1:]...)
		dfs(data, results, append(result, value))
	}
}


