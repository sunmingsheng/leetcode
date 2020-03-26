package main

import "fmt"

//平面上有 n 个点，点的位置用整数坐标表示 points[i] = [xi, yi]。请你计算访问所有这些点需要的最小时间（以秒为单位）。
//
//你可以按照下面的规则在平面上移动：
//
//每一秒沿水平或者竖直方向移动一个单位长度，或者跨过对角线（可以看作在一秒内向水平和竖直方向各移动一个单位长度）。
//必须按照数组中出现的顺序来访问这些点。
//
//输入：points = [[1,1],[3,4],[-1,0]]
//输出：7
//解释：一条最佳的访问路径是： [1,1] -> [2,2] -> [3,3] -> [3,4] -> [2,3] -> [1,2] -> [0,1] -> [-1,0]
//从 [1,1] 到 [3,4] 需要 3 秒
//从 [3,4] 到 [-1,0] 需要 4 秒
//一共需要 7 秒
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-time-visiting-all-points
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	points := [][]int{{1,1},{3,4},{-1,0}}
	fmt.Println(minTimeToVisitAllPoints(points))
}

func minTimeToVisitAllPoints(points [][]int) int {
	results := 0
	for i := 0; i < len(points) - 1; i++ {
		x := points[i][0] - points[i + 1][0]
		if x < 0 {
			x = 0 - x
		}
		y := points[i][1] - points[i + 1][1]
		if y < 0 {
			y = 0 - y
		}
		if x > y {
			results += x
 		} else {
 			results += y
		}
	}
	return results
}
