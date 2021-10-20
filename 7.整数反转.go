package main

import "math"

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 *
 * 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
 * 如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。
 * 假设环境不允许存储 64 位整数（有符号或无符号）。
 *
 * 示例 1：
 * 输入：x = 123
 * 输出：321
 * 示例 2：
 * 输入：x = -123
 * 输出：-321
 *
 * result:
 * Accepted
 * 1032/1032 cases passed (0 ms)
 * Your runtime beats 100 % of golang submissions
 * Your memory usage beats 100 % of golang submissions (2.1 MB)
 */

// @lc code=start
func reverse(x int) int {
	ans := 0
	for x != 0 {
		if ans < math.MinInt32/10 || ans > math.MaxInt32/10 {
			return 0
		}
		// 通过取余进行反转
		ans = ans*10 + x%10
		x /= 10
	}
	return ans
}

// @lc code=end
