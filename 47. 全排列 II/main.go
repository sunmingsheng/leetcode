package main

import "fmt"

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
	temp := make([]int, 0)
	temp = append(temp, 1)
	fmt.Println(temp)
	//fmt.Println(permuteUnique([]int{1,1,2}))
}

func permuteUnique(nums []int) [][]int {
	results := [][]int{}
	result := []int{}
	helper(nums, result, &results)
	return results
}

func helper(nums []int, result []int, results *[][]int) {
	if len(nums) == 0 {
		*results = append(*results, result)
		return
	}
	keys := make(map[int]struct{})
	for key, value := range nums {
		if _, ok := keys[value]; ok {
			continue
		}
		temp := make([]int, key)
		copy(temp, nums[:key])
		temp = append(temp, nums[key+1:]...)
		helper(temp, append(result, value), results)
		keys[value] = struct{}{}
	}
}