package main

import "fmt"

//给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
//
//示例 1:
//
//输入:
//[
//[ 1, 2, 3 ],
//[ 4, 5, 6 ],
//[ 7, 8, 9 ]
//]
//输出: [1,2,3,6,9,8,7,4,5]
//示例 2:
//
//输入:
//[
//[1, 2, 3, 4],
//[5, 6, 7, 8],
//[1,-2,-3,-4]
//[9,10,11,12]
//]
//输出: [1,2,3,4,8,12,11,10,9,5,6,7]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/spiral-matrix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(spiralOrder([][]int{{1},{2},{3}}))
}

func spiralOrder(matrix [][]int) []int {
	var result []int
	m := len(matrix)
	if m == 0 {
		return result
	}
	n := len(matrix[0])
	if n == 0 {
		return result
	}
    up, bottom, left, right := 0, m-1, 0, n-1
    direct := "right"
    for up <= bottom && left <= right {
    	//右
    	if direct == "right" {
    		for i:= left; i <= right; i++ {
    			result = append(result, matrix[up][i])
			}
    		direct = "bottom"
    		up++
    		continue
		}
    	//下
		if direct == "bottom" {
			for i:= up; i <= bottom; i++ {
				result = append(result, matrix[i][right])
			}
			direct = "left"
			right--
			continue
		}
    	//左
		if direct == "left" {
			for i:= right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			direct = "up"
			bottom--
			continue
		}
    	//上
		if direct == "up" {
			for i:= bottom; i >= up; i-- {
				result = append(result, matrix[i][left])
			}
			direct = "right"
			left++
			continue
		}
	}
	return  result
}
