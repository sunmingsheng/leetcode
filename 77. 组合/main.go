package main

import "fmt"

func main() {
	n := 4
	k := 2
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
	dfs(nums, 0, k, result, &results)
	return results
}

func dfs(nums []int, index int, k int,  result []int, results *[][]int) {
	if len(result) == k {
		*results = append(*results, result)
		return
	}
	for i := index; i < len(nums); i++ {
		temp := make([]int, len(result))
		copy(temp, result)
		dfs(nums, i + 1, k, append(temp, nums[i]), results)
	}
}
