package main

import "fmt"

//给定一个未排序的整数数组，找出最长连续序列的长度。
//
//要求算法的时间复杂度为 O(n)。
//
//示例:
//
//输入: [100, 4, 200, 1, 3, 2]
//输出: 4
//解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	s := 0
	for _, value := range nums {
		m[value] = struct{}{}
	}
	for key, _ := range m {
		temp := 1
		data := key
		for {
			if _, ok := m[data + 1]; ok {
				temp += 1
				data += 1
			} else {
				break
			}
		}
		if temp > s {
			s = temp
		}
	}
	return s
}
