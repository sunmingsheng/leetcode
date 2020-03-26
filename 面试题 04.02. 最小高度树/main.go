package main

//给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。
//
//示例:
//给定有序数组: [-10,-3,0,5,9],
//
//一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
//
//0
/// \
//-3   9
///   /
//-10  5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-height-tree-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
    Left *TreeNode
	Right *TreeNode
}

func main() {
	nums := []int{-10,-3,0,5,9}
	sortedArrayToBST(nums)
}

//思路:高度最小的二叉搜索树head元素的val必须是居中的
//每个父节点的值必须是其叶子节点及其自身节点值的居中值
func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return &TreeNode{Val:nums[0]}
	}
	middle := length / 2
	return &TreeNode{
		Val:   nums[middle],
		Left:  sortedArrayToBST(nums[:middle]),
		Right: sortedArrayToBST(nums[middle + 1 :]),
	}
}