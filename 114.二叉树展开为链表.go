package main

/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表(Medium)
 *
 * 给你二叉树的根结点 root ，请你将它展开为一个单链表：
 *  展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
 *  展开后的单链表应该与二叉树 先序遍历 顺序相同。
 *
 * 示例1:
 * 输入：root = [1,2,5,3,4,null,6]
 * 输出：[1,null,2,null,3,null,4,null,5,null,6]
 *
 * result:
 * Accepted
 * 225/225 cases passed (0 ms)
 * Your runtime beats 100 % of golang submissions
 * Your memory usage beats 76.04 % of golang submissions (2.8 MB)
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

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	// 1、左右子树已经被拉平成一条链表
	left := root.Left
	right := root.Right

	// 2、将左子树作为右子树
	root.Left = nil
	root.Right = left

	// 3、将原先的右子树接到当前右子树的末端
	p := root
	for p.Right != nil {
		p = p.Right
	}

	p.Right = right

}

// @lc code=end
