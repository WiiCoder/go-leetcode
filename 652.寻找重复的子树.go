package main

import "fmt"

/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树(Medium)
 *
 * 给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
 * 两棵树重复是指它们具有相同的结构以及相同的结点值。
 *
 * 示例1:
 *       1
 *      / \
 *     2   3
 *    /   / \
 *   4   2   4
 *      /
 *     4
 * 下面是两个重复的子树：
 *     2
 *    /
 *   4
 * 和
 *   4
 *
 * Accepted
 * 176/176 cases passed (16 ms)
 * Your runtime beats 50.68 % of golang submissions
 * Your memory usage beats 19.18 % of golang submissions (12.6 MB)
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

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var key = make(map[string]int)
	var res = []*TreeNode{}
	traverse(root, key, &res)
	return res
}

func traverse(root *TreeNode, key map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return "#"
	}
	left := traverse(root.Left, key, res)
	right := traverse(root.Right, key, res)

	subTree := fmt.Sprintf("%d,%s,%s", root.Val, left, right)

	if v, ok := key[subTree]; !ok {
		key[subTree] = 1
	} else {
		if v == 1 {
			*res = append(*res, root)
		}
		key[subTree]++
	}

	return subTree
}

// @lc code=end
