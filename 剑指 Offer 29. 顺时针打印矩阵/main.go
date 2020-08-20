package main

//输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
//
// 
//
//示例 1：
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//示例 2：
//
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

}

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	length := len(matrix)
	for i := 0; i < length; i++ {
		if i % 2 == 0 {
			for j := 0; j < len(matrix[i]); j++ {
				res = append(res, matrix[i][j])
			}
		} else {
			for j := len(matrix[i]) - 1; j >= 0; j-- {
				res = append(res, matrix[i][j])
			}
		}
	}
	return res
}