package main

import "fmt"

func main() {
	nums := []int{1,2,3}
	fmt.Println(permute(nums))
}

//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
//示例:
//
//输入: [1,2,3]
//输出:
//[
//[1,2,3],
//[1,3,2],
//[2,1,3],
//[2,3,1],
//[3,1,2],
//[3,2,1]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/permutations
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//回溯算法
func permute(nums []int) [][]int {
	results := [][]int{}
	result := []int{}
    back46(nums, result, &results)
	return results
}

func back46(nums, result []int, results *[][]int) {
	if len(nums) <= 0 {
		*results = append(*results, result)
		return
	}
	for key, value := range nums {
		temp := make([]int, key)
		copy(temp, nums[:key])
		temp = append(temp, nums[key+1:]...)
		back46(temp, append(result, value), results)
	}
}

