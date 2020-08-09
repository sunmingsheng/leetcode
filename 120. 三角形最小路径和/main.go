package main

import (
	"fmt"
)

//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//
//相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
//
// 
//
//例如，给定三角形：
//
//[
//[2],
//[3,4],
//[6,5,7],
//[4,1,8,3]
//]
//自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//
// 
//
//说明：
//
//如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/triangle
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(minimumTotal([][]int{{2},{3,4},{6,5,7},{4,1,8,3}}))
}

func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	for j := length - 2; j >= 0 ; j-- {
		for key, _ := range triangle[j] {
			triangle[j][key] += min(triangle[j+1][key], triangle[j+1][key + 1])
		}
	}
	return triangle[0][0]
}

func min(value1, value2 int) int {
	if value1 < value2 {
		return value1
	}
	return value2
}

//func minimumTotal(triangle [][]int) int {
//	length := len(triangle)
//	if length == 0 {
//		return 0
//	}
//	result := math.MaxInt64
//	sum(triangle, 0, 0, 0, &result)
//	return result
//}

//func sum(data[][]int, level int, index int, temp int, result *int) {
//	temp += data[level][index]
//	if level == len(data) - 1 {
//		if temp < *result {
//			*result = temp
//		}
//		return
//	}
//	sum(data, level + 1, index, temp, result)
//	sum(data, level + 1, index + 1, temp, result)
//}