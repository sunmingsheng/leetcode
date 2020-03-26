package main

//给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：
//
//二叉树的根是数组中的最大元素。
//左子树是通过数组中最大值左边部分构造出的最大二叉树。
//右子树是通过数组中最大值右边部分构造出的最大二叉树。
//通过给定的数组构建最大二叉树，并且输出这个树的根节点。
//
// 
//
//示例 ：
//
//输入：[3,2,1,6,0,5]
//输出：返回下面这棵树的根节点：
//
//6
///   \
//3     5
//\    /
//2  0
//\
//1
// 
//
//提示：
//
//给定的数组的大小在 [1, 1000] 之间。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	nums := []int{3,2,1,6,0,5}
	constructMaximumBinaryTree(nums)
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return &TreeNode{
			Val:   nums[0],
		}
	}
	//获取最大值对应的索引
	index := 0
	max := nums[0]
	for key, value := range nums {
		if value > max {
			max = value
			index = key
		}
	}
	return &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
}
