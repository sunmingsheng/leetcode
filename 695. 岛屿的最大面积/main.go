package main

//给定一个包含了一些 0 和 1 的非空二维数组 grid 。
//
//一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
//
//找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)
//
// 
//
//示例 1:
//
//[[0,0,1,0,0,0,0,1,0,0,0,0,0],
//[0,0,0,0,0,0,0,1,1,1,0,0,0],
//[0,1,1,0,1,0,0,0,0,0,0,0,0],
//[0,1,0,0,1,1,0,0,1,0,1,0,0],
//[0,1,0,0,1,1,0,0,1,1,1,0,0],
//[0,0,0,0,0,0,0,0,0,0,1,0,0],
//[0,0,0,0,0,0,0,1,1,1,0,0,0],
//[0,0,0,0,0,0,0,1,1,0,0,0,0]]
//对于上面这个给定矩阵应返回 6。注意答案不应该是 11 ，因为岛屿只能包含水平或垂直的四个方向的 1 。
//
//示例 2:
//
//[[0,0,0,0,0,0,0,0]]
//对于上面这个给定的矩阵, 返回 0
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/max-area-of-island
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

}

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	if m <= 0 {
		return 0
	}
	n := len(grid[0])
	if n <= 0 {
		return 0
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := dfs(grid, m, n, i, j)
				if area > res {
					res = area
				}
			}
		}
	}
	return res
}

func dfs(grid [][]int, m, n, i, j int) int {
	if i < 0 || i > m - 1 || j < 0 || j > n - 1 || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	count := 1
	count += dfs(grid, m, n, i-1, j)
	count += dfs(grid, m, n, i+1, j)
	count += dfs(grid, m, n, i, j-1)
	count += dfs(grid, m, n, i, j+1)
	return count
}