package main

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
//岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
//
//此外，你可以假设该网格的四条边均被水包围。
//
// 
//
//示例 1:
//
//输入:
//[
//['1','1','1','1','0'],
//['1','1','0','1','0'],
//['1','1','0','0','0'],
//['0','0','0','0','0']
//]
//输出: 1
//示例 2:
//
//输入:
//[
//['1','1','0','0','0'],
//['1','1','0','0','0'],
//['0','0','1','0','0'],
//['0','0','0','1','1']
//]
//输出: 3
//解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/number-of-islands
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

}


func numIslands(grid [][]byte) int {
	m := len(grid)
	if m <= 0 {
		return 0
	}
	n := len(grid[0])
	if n <= 0 {
		return 0
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, m, n, i, j)
			}
		}
	}
	return count
}

func dfs (grid [][]byte, m, n, i, j int) {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, m, n,  i+1, j)
	dfs(grid, m, n,  i-1, j)
	dfs(grid, m, n,  i, j+1)
	dfs(grid, m, n,  i, j-1)
}
