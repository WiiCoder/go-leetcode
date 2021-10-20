package main

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数(Easy)
 *
 * 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
 * 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
 *
 * 示例 1：
 * 输入：x = 121
 * 输出：true
 *
 * 示例 2：
 * 输入：x = -121
 * 输出：false
 * 解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
 *
 * 示例 3：
 * 输入：x = 10
 * 输出：false
 * 解释：从右向左读, 为 01 。因此它不是一个回文数。
 *
 * result:
 * Accepted
 * 11510/11510 cases passed (32 ms)
 * Your runtime beats 7.14 % of golang submissions
 * Your memory usage beats 91.25 % of golang submissions (4.9 MB)
 */

// @lc code=start
func isPalindrome(x int) bool {
	ans := 0
	res := x
	// 对X进行反转
	for x != 0 && x > 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	// 反转后的结果与x相等则说明x是回文数
	return ans == res
}

// @lc code=end
