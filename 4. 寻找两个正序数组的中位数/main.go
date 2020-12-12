package main

import "fmt"

//给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
//
//请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
//你可以假设 nums1 和 nums2 不会同时为空。
//
// 
//
//示例 1:
//
//nums1 = [1, 3]
//nums2 = [2]
//
//则中位数是 2.0
//示例 2:
//
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//则中位数是 (2 + 3)/2 = 2.5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(findMedianSortedArrays([]int{1,3}, []int{2,4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length1 := len(nums1)
	length2 := len(nums2)
	arr := []int{}
	i := 0
	j := 0
	for {
		if i >= length1 || j >= length2 {
			break
		}
		if nums1[i] <= nums2[j] {
			arr = append(arr,nums1[i])
			i++
		} else {
			arr = append(arr,nums2[j])
			j++
		}
	}
	if i < length1 {
		arr = append(arr, nums1[i:]...)
	}
	if j < length2 {
		arr = append(arr, nums2[j:]...)
	}
	if (length1 + length2) % 2  == 0 {
		return (float64(arr[(length1 + length2) / 2 - 1]) + float64(arr[(length1 + length2) / 2])) / 2
	} else {
		return float64(arr[(length1 + length2) / 2])
	}
}
