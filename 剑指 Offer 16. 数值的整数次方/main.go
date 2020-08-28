package main

//实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。
//
// 
//
//示例 1:
//
//输入: 2.00000, 10
//输出: 1024.00000
//示例 2:
//
//输入: 2.10000, 3
//输出: 9.26100
//示例 3:
//
//输入: 2.00000, -2
//输出: 0.25000
//解释: 2-2 = 1/22 = 1/4 = 0.25
// 
//
//说明:
//
//-100.0 < x < 100.0
//n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


func myPow(x float64, n int) float64 {
	m := make(map[int]float64)
	return myPow_(x, n, m)
}

func myPow_(x float64, n int, m map[int]float64) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1 / x
	}
	data := 0.00
	if result, ok := m[n]; ok {
		return result
	}
	if n % 2 == 1 || n % 2 == -1 {
		//奇数
		if n < 0 {
			data = myPow_(x,n/2,m) * myPow_(x,n/2,m) / x
		} else {
			data = myPow_(x,n/2,m) * myPow_(x,n/2,m) * x
		}
	} else {
		//偶数
		data =  myPow_(x,n/2,m) * myPow_(x,n/2,m)
	}
	m[n] = data
	return data
}