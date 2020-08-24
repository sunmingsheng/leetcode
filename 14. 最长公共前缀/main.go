package main

import (
	"fmt"
	"math"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//示例 1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//说明:
//
//所有输入只包含小写字母 a-z 。
//
//通过次数341,416提交次数88
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-common-prefix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main()  {
	fmt.Println(longestCommonPrefix([]string{"a","b",""}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLength := math.MaxInt64
	minStr := ""
	for _, str := range strs {
		if len(str) < minLength {
			minLength = len(str)
			minStr = str
		}
	}
	for i := minLength; i >= 0; i-- {
		temp := minStr[0:i]
		flag := true
		for _, value := range strs {
			if value[0:i] != temp {
				flag = false
				break
			}
		}
		if flag {
			return temp
		}
	}
	return ""
}
