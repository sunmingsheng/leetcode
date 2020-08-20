package main

import "fmt"

//0,1,,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字。求出这个圆圈里剩下的最后一个数字。
//
//例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。
//
// 
//
//示例 1：
//
//输入: n = 5, m = 3
//输出: 3
//示例 2：
//
//输入: n = 10, m = 17
//输出: 2
// 
//
//限制：
//
//1 <= n <= 10^5
//1 <= m <= 10^6
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(lastRemaining(10, 17))
}

//n = 5 m = 3
//0 1 2 3 4 idx = 2
//0 1 3 4
//环形链表也可以实现就是时间复杂度高
func lastRemaining(n int, m int) int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}
	idx := 0
	for n > 1 {
		idx = (idx + m - 1) % n //取余算法很关键
		data = append(data[0:idx], data[idx+1:]...)
		n--
	}
	return data[0]
}

