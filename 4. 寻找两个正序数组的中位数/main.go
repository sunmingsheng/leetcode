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
	fmt.Println(findMedianSortedArrays([]int{1,3}, []int{2}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	res  := make([]int, len1 + len2)
	if len1 == 0 || len2 == 0 {
		res = append(nums1, nums2...)
	} else {
		index := 0
		i, j := 0, 0
		for ; i < len1; i++ {
			temp1 := nums1[i]
			for ; j < len2; j++ {
				temp2 := nums2[j]
				if temp1 > temp2 {
					res[index] = temp2
					index += 1
				} else {
					break
				}
			}
			res[index] = temp1
			index += 1
		}
		for i < len1 {
			res[index] = nums1[i]
			i ++
			index ++
		}
		for j < len2 {
			res[index] = nums2[j]
			j ++
			index ++
		}
	}
	if (len1 + len2) % 2 == 0 {
		//偶数
		temp1 := float64(res[(len1 + len2) / 2])
		temp2 := float64(res[(len1 + len2) / 2 - 1])
		return (temp1 + temp2) / 2
	} else {
		//
		index := (len1 + len2) / 2
		return float64(res[index])
	}
}
