package main

/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针(Medium)
 *
 * 给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。
 * 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
 * 初始状态下，所有 next 指针都被设置为 NULL。
 *
 * result:
 * Accepted
 * 58/58 cases passed (8 ms)
 * Your runtime beats 26.9 % of golang submissions
 * Your memory usage beats 77.16 % of golang submissions (6.3 MB)
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

func connectTwoNode(one *Node, two *Node) {
	if one == nil || two == nil {
		return
	}
	one.Next = two

	// 建立子节点连接
	connectTwoNode(one.Left, one.Right)
	connectTwoNode(two.Left, two.Right)

	// 建立跨节点连接
	connectTwoNode(one.Right, two.Left)
}

// @lc code=end
