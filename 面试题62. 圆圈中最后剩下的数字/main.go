package main

import (
	"fmt"
)

func main() {
	fmt.Println(lastRemaining(70866, 116922))
}

//0 1 2 3 4
func lastRemaining(n int, m int) int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	result := 0
	index := 1
	point := 0 //指向数组
	for {
		if len(s) == 1 {
			result = s[0]
			break
		}
		if index == m {
			temp := make([]int, point)
			copy(temp, s[:point])
			s = append(temp, s[point+1:]...)
			index = 1
			point--
		} else {
			index++
		}
		if point + 1 >= len(s) {
			point = 0
		} else {
			point++
		}
	}
	return result
}