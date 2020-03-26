package main

import (
	"fmt"
	"sort"
)

//幂集。编写一种方法，返回某集合的所有子集。集合中不包含重复的元素。
//
//说明：解集不能包含重复的子集。
//
//示例:
//
//输入： nums = [1,2,3]
//输出：
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
//链接：https://leetcode-cn.com/problems/power-set-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	results := subsets([]int{9,0,3,5,7})
	for _, value := range results {
		fmt.Println(value)
	}
}

func subsets(nums []int) [][]int {
	results := [][]int{}
	result  := []int{}
	back4(nums, result, &results)
	return results
}

func back4(nums []int, result []int, results *[][]int) {
	sort.Ints(result)
    if len(nums) <= 0 {
    	*results = append(*results, result)
    	return
	} else {
		*results = append(*results, result)
		for key, value := range nums {
			tempResult := append(result, value)
			temp := make([]int, len(nums) - key - 1)
			copy(temp, nums[key + 1:])
			back4(temp, tempResult, results)
		}
	}
}