package main

/*
 * @lc app=leetcode.cn id=700 lang=golang
 *
 * [700] 二叉搜索树中的搜索
 * 给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。
 *
 * 示例:
 * 给定二叉搜索树:
 *
 *         4
 *        / \
 *       2  7
 *      / \
 *     1  3
 *
 * 和值: 2
 * 返回如下子树:
 *   	2
 *     / \
 *    1  3
 * 在上述示例中，如果要找的值是 5，但因为没有节点值为 5，我们应该返回 NULL。
 *
 * result:
 * Accepted
 * 36/36 cases passed (24 ms)
 * Your runtime beats 59.77 % of golang submissions
 * Your memory usage beats 87.81 % of golang submissions (7.2 MB)
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

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}

	if root.Val < val {
		return searchBST(root.Right, val)
	}

	return root
}

// @lc code=end
