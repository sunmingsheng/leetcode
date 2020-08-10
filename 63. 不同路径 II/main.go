package main

import "fmt"

//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
//现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//
//
//
//网格中的障碍物和空位置分别用 1 和 0 来表示。
//
//说明：m 和 n 的值均不超过 100。
//
//示例 1:
//
//输入:
//[
//  [0,0,0],
//  [0,1,0],
//  [0,0,0]
//]
//输出: 2
//解释:
//3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/unique-paths-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{1}}))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	xx := len(obstacleGrid)
	if xx <= 0 {
		return 0
	}
	yy := len(obstacleGrid[0])
	if yy <= 0 {
		return 0
	}
	dp := make([][]int, xx)
	for x, value := range obstacleGrid {
		dp[x] = make([]int, yy)
		for y, v := range value {
			if v == 1 {
				dp[x][y] = 0
			} else if x == 0 && y == 0 {
				dp[x][y] = 1
			} else if x == 0 || y == 0 {
				if x == 0 {
					dp[x][y] = dp[x][y-1]
				}
				if y == 0 {
					dp[x][y] = dp[x-1][y]
				}
			} else {
				dp[x][y] = dp[x-1][y] + dp[x][y-1]
			}
		}
	}
	return dp[xx-1][yy-1]
}

//func uniquePathsWithObstacles(obstacleGrid [][]int) int {
//	x := len(obstacleGrid)
//	if x <= 0 {
//		return 0
//	}
//	y := len(obstacleGrid[0])
//	if y <= 0 {
//		return 0
//	}
//	c := make(map[int]map[int]int)
//	return travel(obstacleGrid, c, x, y, 0, 0)
//}
//
//func travel(data [][]int, c map[int]map[int]int, x, y, nowX, nowY int) int {
//	if nowX >= x || nowY >= y {
//		return 0
//	}
//	if data[nowX][nowY] == 1 {
//		return 0
//	}
//	if nowX == x-1 && nowY == y-1 {
//		return 1
//	}
//	if _, ok := c[nowX]; !ok {
//		data := travel(data, c, x, y, nowX + 1, nowY) + travel(data, c, x, y, nowX, nowY + 1)
//		c[nowX] = make(map[int]int)
//		c[nowX][nowY] = data
//		return data
//	} else {
//		if cache, ok := c[nowX][nowY]; !ok {
//			data := travel(data, c, x, y, nowX + 1, nowY) + travel(data, c, x, y, nowX, nowY + 1)
//			c[nowX][nowY] = data
//			return data
//		} else {
//			return cache
//		}
//	}
//}
