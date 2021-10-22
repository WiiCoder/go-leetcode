package main

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加(Medium)
 * 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
 * 请你将两个数相加，并以相同形式返回一个表示和的链表。
 * 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 * 示例 1：
 * 输入：l1 = [2,4,3], l2 = [5,6,4]
 * 输出：[7,0,8]
 * 解释：342 + 465 = 807.
 *
 * 示例 2：
 * 输入：l1 = [0], l2 = [0]
 * 输出：[0]
 *
 * result:
 * Accepted
 * 1568/1568 cases passed (12 ms)
 * Your runtime beats 57.11 % of golang submissions
 * Your memory usage beats 100 % of golang submissions (4.4 MB)
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h := l1
	z := 0
	x := h
	for !(l1 == nil && l2 == nil) {
		a := 0
		b := 0
		if l1 != nil {
			a = l1.Val
		}
		if l2 != nil {
			b = l2.Val
		}
		// 取模
		x.Val = (a + b + z) % 10
		z = (a + b + z) / 10
		// 判空跳出循环
		if (l2 == nil || l1 == nil) && z == 0 {
			break
		}
		// 修改节点指向
		if x.Next == nil && l2 != nil {
			l1 = nil
			x.Next = l2.Next
		}
		if x.Next == nil && z != 0 {
			x.Next = &ListNode{Val: z}
			break
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		// 指向下一个节点
		x = x.Next
	}
	return h
}

// @lc code=end
