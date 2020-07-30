package main

import (
	"fmt"
	"sort"
)

//给出一个区间的集合，请合并所有重叠的区间。
//
//示例 1:
//
//输入: [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例 2:
//
//输入: [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-intervals
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	data := [][]int{{1,3},{2,6},{8,10},{15,18}}
	fmt.Println(merge(data))
}


func merge(intervals [][]int) [][]int {
	length := len(intervals)
	if length <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] > intervals[j][0] {
			return false
		}
		return true
	})
	data := [][]int{}
	for i := 0; i < length; i++ {
		if i == length - 1 {
			data = append(data, intervals[i])
			break
		}
		j := i + 1
		if intervals[i][1] >= intervals[j][0] {
			intervals[j][0] = intervals[i][0]
			if intervals[i][1] >= intervals[j][1] {
				intervals[j][1] = intervals[i][1]
			} else {
				intervals[j][1] = intervals[j][1]
			}
		} else {
			data = append(data, intervals[i])
		}
	}
	return data
}
