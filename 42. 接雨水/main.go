package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	//1 + 3 + 1
}

func trap(data []int) int {
	res := 0
	length := len(data)
	if length <= 2 {
		return res
	}
	maxIndex := -1
	maxHeight := 0
	for key, value := range data {
		if value > maxHeight {
			maxIndex = key
			maxHeight = value
		}
	}
	//左边
	maxLeftHight := data[0]
	for i := 1; i < maxIndex; i++ {
		if data[i] >= maxLeftHight {
			maxLeftHight = data[i]
		} else {
			tmp := maxLeftHight - data[i]
			res += tmp
		}
	}
	//右边
	maxRightHight := data[length-1]
	for i := length-2; i > maxIndex; i-- {
		if data[i] >= maxRightHight {
			maxRightHight = data[i]
		} else {
			tmp := maxRightHight - data[i]
			res += tmp
		}
	}
	return res
}

