package main

import (
	"fmt"
	"math"
)

//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
//
// 
//
//示例：
//
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
// 
//
//提示：
//
//3 <= nums.length <= 10^3
//-10^3 <= nums[i] <= 10^3
//-10^4 <= target <= 10^4
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/3sum-closest
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(threeSumClosest([]int{-1,2,1,4}, 1))
}


func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	if length < 3 {
		return 0
	}
	if length == 3 {
		return sum(nums[0], nums[1], nums[2])
	}
	res := 0
	diff := math.MaxInt64
	for i := 0; i < length - 2; i++ {
		for j := i + 1; j < length - 1; j++ {
			for z := j + 1; z < length; z++ {
				temp := sum(nums[i], nums[j], nums[z])
				if temp == target {
					return target
				}
				data := 0
				if temp > target {
					data = temp - target
				} else {
					data = target - temp
				}
				if data < diff {
					diff = data
					res = temp
				}
			}
		}
	}
	return  res
}

func sum(v1, v2, v3 int) int {
	return v1 + v2 + v3
}



