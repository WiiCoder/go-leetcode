package main

/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
 * 给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
 * 提醒一下，二叉搜索树满足下列约束条件：
 *  - 节点的左子树仅包含键 小于 节点键的节点。
 *  - 节点的右子树仅包含键 大于 节点键的节点。
 *  - 左右子树也必须是二叉搜索树。
 *
 * result:
 * Accepted
 * 215/215 cases passed (12 ms)
 * Your runtime beats 82.4 % of golang submissions
 * Your memory usage beats 65.57 % of golang submissions (6.8 MB)
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
func convertBST(root *TreeNode) *TreeNode {
	res := 0
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Right)
		res += root.Val
		root.Val = res
		traverse(root.Left)
	}
	traverse(root)
	return root
}

// @lc code=end
