package main

import "fmt"

//输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。
//
// 
//
//示例 1：
//
//输入：nums = [2,7,11,15], target = 9
//输出：[2,7] 或者 [7,2]
//示例 2：
//
//输入：nums = [10,26,30,31,47,60], target = 40
//输出：[10,30] 或者 [30,10]
// 
//
//限制：
//
//1 <= nums.length <= 10^5
//1 <= nums[i] <= 10^6
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(twoSum([]int{14,15,16,22,53,60}, 76))
}

func twoSum(nums []int, target int) []int {
	length := len(nums)
	if length < 2 {
		return []int{}
	}
	left := 0
	right := length - 1
	for i := 0; i < length; i++ {
		if nums[i] >= target {
			right --
		}
	}
	for left < right {
		temp := nums[left] + nums[right]
		if temp == target {
			return []int{nums[left], nums[right]}
		}
		if temp > target {
			right --
		} else {
			left ++
		}
	}
	return []int{}
}

