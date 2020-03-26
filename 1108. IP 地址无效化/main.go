package main

import (
	"fmt"
)

//给你一个有效的 IPv4 地址 address，返回这个 IP 地址的无效化版本。
//
//所谓无效化 IP 地址，其实就是用 "[.]" 代替了每个 "."。
//
// 
//
//示例 1：
//
//输入：address = "1.1.1.1"
//输出："1[.]1[.]1[.]1"
//示例 2：
//
//输入：address = "255.100.50.0"
//输出："255[.]100[.]50[.]0"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/defanging-an-ip-address
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}

func defangIPaddr(address string) string {
	r := []rune(address)
	s := ""
	for _, value := range r {
		temp := string(value)
		if temp == "." {
			s += "[.]"
		} else {
			s += temp
		}
	}
	return s
}
