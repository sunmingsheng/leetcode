package main

import (
	"fmt"
)

//请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。
//
//[["a","b","c","e"],
//["s","f","c","s"],
//["a","d","e","e"]]
//
//但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
//
// 
//
//示例 1：
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
//输出：true
//示例 2：
//
//输入：board = [["a","b"],["c","d"]], word = "abcd"
//输出：false
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCESEEEFS"
	fmt.Println(exist(board,word))
}

//[
// ["A","B","C","E"],
// ["S","F","E","S"],
// ["A","D","E","E"]
// ]
//"ABCESEEEFS"

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return false
	}
	yLen := len(board)
	xLen := len(board[0])
	for i := 0 ; i < yLen; i++ {
		for j := 0; j < xLen; j++ {
			if board[i][j] == word[0] {
				m := make([][]int, yLen)
				for key, _ := range m {
					m[key] = make([]int, xLen)
				}
				m[i][j] = 1
				if helper(board, i, j, m, word[1:]) {
					return true
				}
			}
		}
	}
	return false
}

func helper(board [][]byte, i, j int, m [][]int, word string) bool {
	if len(word) == 0 {
		return true
	}
	//往四个方向走
	//往左
	if j >= 1 && m[i][j-1] == 0 && board[i][j-1] == word[0] {
		m[i][j-1] = 1
		if helper(board, i, j-1 , m, word[1:]) {
			return true
		}
	}
	//往右
	if j < len(board[0]) - 1 && m[i][j+1] == 0 && board[i][j+1] == word[0] {
		m[i][j+1] = 1
		if helper(board, i, j+1 , m, word[1:]) {
			return true
		}
	}
	//向上
	if i >= 1 && m[i-1][j] == 0 && board[i-1][j] == word[0] {
		m[i-1][j] = 1
		if helper(board, i-1, j , m, word[1:]) {
			return true
		}
	}
	//往下
	if i < len(board) - 1 && m[i+1][j] == 0 && board[i+1][j] == word[0] {
		m[i+1][j] = 1
		if helper(board, i+1, j, m, word[1:]) {
			return true
		}
	}
	m[i][j] = 0
	return false
}