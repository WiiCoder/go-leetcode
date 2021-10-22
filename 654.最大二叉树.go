package main

/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树(Medium)
 *
 * 给定一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：
 *
 * 二叉树的根是数组 nums 中的最大元素。
 * 左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
 * 右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
 * 返回有给定数组 nums 构建的 最大二叉树 。
 *
 * 示例1:
 * 输入：nums = [3,2,1,6,0,5]
 * 输出：[6,3,5,null,2,0,null,null,1]
 * 解释：递归调用如下所示：
 * - [3,2,1,6,0,5] 中的最大值是 6 ，左边部分是 [3,2,1] ，右边部分是 [0,5] 。
 *     - [3,2,1] 中的最大值是 3 ，左边部分是 [] ，右边部分是 [2,1] 。
 *         - 空数组，无子节点。
 *         - [2,1] 中的最大值是 2 ，左边部分是 [] ，右边部分是 [1] 。
 *             - 空数组，无子节点。
 *             - 只有一个元素，所以子节点是一个值为 1 的节点。
 *     - [0,5] 中的最大值是 5 ，左边部分是 [0] ，右边部分是 [] 。
 *         - 只有一个元素，所以子节点是一个值为 0 的节点。
 *         - 空数组，无子节点。
 *
 * result:
 * Accepted
 * 107/107 cases passed (16 ms)
 * Your runtime beats 84.95 % of golang submissions
 * Your memory usage beats 93.93 % of golang submissions (7 MB)
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	max := 0
	maxIndex := 0
	// 取最大值
	for i, v := range nums {
		if v > max {
			max = v
			maxIndex = i
		}
	}
	// 建立根节点
	root := &TreeNode{
		Val: max,
	}
	// 构建左树
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	// 构建右树
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])

	return root
}

// @lc code=end
