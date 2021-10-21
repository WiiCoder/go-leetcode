package main

/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树(Easy)
 *
 * 翻转一棵二叉树。
 *
 * 示例：
 *  输入：
 *      4
 *    /   \
 *   2     7
 *  / \   / \
 * 1   3 6   9
 *
 * 输出：
 *      4
 *    /   \
 *   7     2
 *  / \   / \
 * 9   6 3   1
 *
 * result:
 * Accepted
 * 77/77 cases passed (0 ms)
 * Your runtime beats 100 % of golang submissions
 * Your memory usage beats 62.43 % of golang submissions (2.1 MB)
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

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 子节点左右交换
	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Left, root.Right = right, left

	return root
}

// @lc code=end
