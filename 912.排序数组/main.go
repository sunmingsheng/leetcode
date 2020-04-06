package main

import (
	"fmt"
	"sort"
)

//给你一个整数数组 nums，将该数组升序排列。
//
// 
//
//示例 1：
//
//输入：nums = [5,2,3,1]
//输出：[1,2,3,5]
//示例 2：
//
//输入：nums = [5,1,1,2,0,0]
//输出：[0,0,1,1,2,5]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sort-an-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(sortArray([]int{1,2,3,1}))
}

func sortArray(nums []int) []int {
	sort.Ints(nums)
	return nums
}

//常用排序算法的总结