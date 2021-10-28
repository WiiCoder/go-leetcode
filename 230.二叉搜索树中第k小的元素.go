package main

/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * 二叉搜索树特点
 * 1、对于 BST 的每一个节点node，左子树节点的值都比node的值要小，右子树节点的值都比node的值大。
 * 2、对于 BST 的每一个节点node，它的左侧子树和右侧子树都是 BST。
 *
 * [230] 二叉搜索树中第K小的元素(Medium)
 * 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
 *
 * 示例 1：
 * 输入：root = [3,1,4,null,2], k = 1
 * 输出：1
 *
 * 示例 2：
 * 输入：root = [5,3,6,2,4,null,null,1], k = 3
 * 输出：3
 *
 * result:
 * Accepted
 * 93/93 cases passed (4 ms)
 * Your runtime beats 99.7 % of golang submissions
 * Your memory usage beats 76.41 % of golang submissions (6.4 MB)
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

func kthSmallest(root *TreeNode, k int) int {
	res := 0
	rank := 0
	var mid func(root *TreeNode, k int)
	mid = func(root *TreeNode, k int) {
		if root == nil {
			return
		}
		mid(root.Left, k)
		rank++
		if k == rank {
			res = root.Val
			return
		}
		mid(root.Right, k)

	}
	mid(root, k)
	return res
}

// @lc code=end
