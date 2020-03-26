package main

import "fmt"

//给你两个整数数组 nums 和 index。你需要按照以下规则创建目标数组：
//
//目标数组 target 最初为空。
//按从左到右的顺序依次读取 nums[i] 和 index[i]，在 target 数组中的下标 index[i] 处插入值 nums[i] 。
//重复上一步，直到在 nums 和 index 中都没有要读取的元素。
//请你返回目标数组。
//
//题目保证数字插入位置总是存在。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/create-target-array-in-the-given-order
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	n1 := []int{1,2,3,4,0}
	n2 := []int{0,1,2,3,0}
	fmt.Println(createTargetArray(n1, n2))
}

func createTargetArray(nums []int, index []int) []int {
	targets := []int{}
	for key, value := range nums {
		i := index[key]
		if len(targets) <= i {
			//追加
			targets = append(targets, value)
		} else {
			//插入
            leftS := targets[:i]
            rightS := targets[i:]
            temp := make([]int, len(leftS))
            copy(temp, leftS)
            temp = append(temp, value)
            temp = append(temp, rightS...)
            targets = temp
		}
	}
	return targets
}
