package main

import (
	"fmt"
	"sort"
)

//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
//说明：解集不能包含重复的子集。
//
//示例:
//
//输入: nums = [1,2,3]
//输出:
//[
//[3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/subsets
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(subsets([]int{1,2,3}))
}

func subsets(nums []int) [][]int {
	results := [][]int{}
	result := []int{}
	sort.Ints(nums)
	dfs(nums, result ,&results,0)
	return results
}

//78-----todo
func dfs(nums []int, result []int, results *[][]int, index int) {
	if index == len(nums) {
		*results = append(*results, result)
		return
	}
	dfs(nums, result, results, index + 1)
	for i := index; i < len(nums); i++ {
		dfs(nums, append(result, nums[i]), results, i + 1)
	}
}

