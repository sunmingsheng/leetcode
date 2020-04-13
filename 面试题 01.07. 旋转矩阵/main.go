package main

import "fmt"

//给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。
//
//不占用额外内存空间能否做到？
//
// 
//示例 1:
//
//给定 matrix =
//[
//[1,2,3],
//[4,5,6],
//[7,8,9]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//[7,4,1],
//[8,5,2],
//[9,6,3]
//]
//示例 2:
//
//给定 matrix =
//[
//[ 5, 1, 9,11],
//[ 2, 4, 8,10],
//[13, 3, 6, 7],
//[15,14,12,16]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//[15,13, 2, 5],
//[14, 3, 4, 1],
//[12, 6, 8, 9],
//[16, 7,10,11]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/rotate-matrix-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	matrix := make([][]int,3,3)
	matrix[0] = []int{1,2,3}
	matrix[1] = []int{4,5,6}
	matrix[2] = []int{7,8,9}
	rotate(matrix)
}

func rotate(matrix [][]int)  {
	length := len(matrix)
	results := make([][]int, length, length)
	for key, _ := range results {
		results[key] = make([]int, length, length)
	}
	for j := 0; j < length; j++ {
		for i := 0; i < length; i++ {
			results[i][j] = matrix[length-j-1][i]
		}
	}
	matrix = results
	fmt.Println(results)
}
