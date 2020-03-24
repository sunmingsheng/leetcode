package main

import (
	"fmt"
	"math"
)

//找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
//
//说明：
//
//所有数字都是正整数。
//解集不能包含重复的组合。 
//示例 1:
//
//输入: k = 3, n = 7
//输出: [[1,2,4]]
//示例 2:
//
//输入: k = 3, n = 9
//输出: [[1,2,6], [1,3,5], [2,3,4]]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/combination-sum-iii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	results := combinationSum3(2, 18)
	for _, value := range results {
		fmt.Println(value)
	}
}

func combinationSum3(k int, n int) [][]int {
	//[1,n),n<=9,有序的集合
	results := [][]int{}
	nums := make([]int, int(math.Min(float64(n),float64(10)))-1)
	for i := 1; i < int(math.Min(float64(n),float64(10))); i++ {
		nums[i - 1] = i
	}
	result := []int{}
	back216(nums, k, n, result, &results)
	return results
}

func back216(nums []int, k int, n int, result []int, results *[][]int) {
	if len(result) == k {
		//求和
		sumResult := 0
		for _, value := range result {
			sumResult += value
		}
		if sumResult == n {
			*results = append(*results, result)
		}
		return
	}
	for key, value := range nums {
		temp := make([]int, len(nums) - key - 1)
		copy(temp, nums[key+1:])
		tempResult := make([]int, len(result))
		copy(tempResult, result)
		tempResult = append(tempResult, value)
		back216(temp, k, n, tempResult, results)
	}
}