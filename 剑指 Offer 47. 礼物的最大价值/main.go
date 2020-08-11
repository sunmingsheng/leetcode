package main

import "fmt"

//在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
//
//示例 1:
//
//输入:
//[
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
//]
//输出: 12
//解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(maxValue([][]int{{1,3,1},{1,5,1},{4,2,1}}))
}

func maxValue(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0 ; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
				continue
			}
			if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j - 1]
				continue
			}
			if j == 0 {
				dp[i][j] = grid[i][j] + dp[i - 1][j]
				continue
			}
			if dp[i - 1][j] > dp[i][j - 1] {
				dp[i][j] = grid[i][j] + dp[i - 1][j]
			} else {
				dp[i][j] = grid[i][j] + dp[i][j - 1]
			}
		}
	}
	return dp[m-1][n-1]
}
