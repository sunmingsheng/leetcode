package main

import "fmt"

//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//
// 
//
//示例 1：
//
//输入：[1,2,3,1]
//输出：4
//解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//示例 2：
//
//输入：[2,7,9,3,1]
//输出：12
//解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/house-robber
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(rob([]int{2,7,9,3,1}))
}

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	dp := make([]int, length)
	dp[0] = nums[0]
	if nums[0] > nums[1] {
		dp[1] = nums[0]
	} else {
		dp[1] = nums[1]
	}
	for i := 2; i < length; i++ {
		if dp[i-1] > dp[i-2] + nums[i] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-2] + nums[i]
		}
	}
	return dp[length - 1]
}

//func rob(nums []int) int {
//	l := 0
//	r := len(nums) - 1
//	m := make(map[int]int)
//	return rob_(nums, l, r, m)
//}
//
//func rob_(nums []int, l, r int, m map[int]int) int {
//	length := r - l
//	if length < 0 {
//		return 0
//	}
//	if length == 0 {
//		return nums[l]
//	}
//	if length == 1 {
//		if nums[l] > nums[r] {
//			return nums[l]
//		}
//		return nums[r]
//	}
//	if value, ok := m[l]; ok {
//		return value
//	}
//	max := 0
//	for i := l; i <= r; i++ {
//		temp := nums[i] + rob_(nums, i + 2, r, m)
//		if temp > max {
//			max = temp
//		}
//	}
//	m[l] = max
//	return max
//}