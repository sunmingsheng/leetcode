package main

import (
	"fmt"
	"sort"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
// 
//
//示例：
//
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//[-1, 0, 1],
//[-1, -1, 2]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/3sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	results := [][]int{}
	sort.Ints(nums)
	length := len(nums)
	twos := make(map[int]map[int]struct{})
	for i := 0; i < length - 2; i++ {
		z := length - 1
		for j := i+1; j < length - 1; j++ {
			two := nums[i] + nums[j]
			for z > j {
				three := two + nums[z]
				if three < 0 {
					break
				}
				if three == 0 {
					if _, ok := twos[nums[i]]; !ok {
						twos[nums[i]] = make(map[int]struct{})
					}
					if _, ok := twos[nums[i]][nums[j]]; !ok {
						twos[nums[i]][nums[j]] = struct{}{}
						results = append(results, []int{nums[i], nums[j], nums[z]})
					}
					if nums[z - 1] != nums[z] {
						break
					}
				}
				z--
			}
		}
	}
	return results
}