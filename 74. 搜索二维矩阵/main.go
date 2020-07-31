package main

import "fmt"

//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//每行中的整数从左到右按升序排列。
//每行的第一个整数大于前一行的最后一个整数。
//示例 1:
//
//输入:
//matrix = [
//[1,   3,  5,  7],
//[10, 11, 16, 20],
//[23, 30, 34, 50]
//]
//target = 3
//输出: true
//示例 2:
//
//输入:
//matrix = [
//[1,   3,  5,  7],
//[10, 11, 16, 20],
//[23, 30, 34, 50]
//]
//target = 13
//输出: false
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/search-a-2d-matrix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	matrix := [][]int {
		{1,   3,  5,  7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(matrix, 50))
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m <= 0 {
		return false
	}
	n := len(matrix[0])
	if n <= 0 {
		return false
	}
	for _, outData := range matrix {
		if outData[0] > target {
			return false
		}
		if outData[n-1] < target {
			continue
		}
		//二分查找
		return search(outData, target)
	}
	return false
}

func search(nums []int, target int) bool {
	length := len(nums)
	if length == 1 {
		if nums[0] == target {
			return true
		}
		return false
	}
	mid := length / 2
	if nums[mid] == target {
		return true
	}
	if nums[mid] > target {
		return search(nums[0:mid], target)
	}
	return search(nums[mid:], target)
}
