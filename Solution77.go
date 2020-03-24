package main

import "fmt"

func main() {
	n := 5
	k := 4
	results := combine(n, k)
	for _, value := range results {
		fmt.Println(value)
	}
}

func combine(n int, k int) [][]int {
	if k <= 0 {
		return [][]int{}
	}
	//生成切片
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i - 1] = i
	}
	results := [][]int{}
	result := []int{}
	//回溯数据
	back77(nums, k, result, &results)
	return results
}

func back77(nums []int, n int, result []int, results *[][]int) {
	if len(result) == n {
		*results = append(*results, result)
		return
	}
	for key, value := range nums {
		temp := make([]int, len(nums) - key - 1)
		copy(temp, nums[key + 1:])
		tempResult := make([]int, len(result))
		copy(tempResult, result)
		tempResult = append(tempResult, value)
		back77(temp, n, tempResult, results)
	}
}