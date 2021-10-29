package main

/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
 * 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
 * 一般来说，删除节点可分为两个步骤：
 *  1.首先找到需要删除的节点；
 *  2.如果找到了，删除它。
 *
 *
 * result:
 * Accepted
 * 91/91 cases passed (12 ms)
 * Your runtime beats 86.9 % of golang submissions
 * Your memory usage beats 64.35 % of golang submissions (6.9 MB)
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		minNode := getMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func getMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// @lc code=end
