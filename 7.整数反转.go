package main

import "math"

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	ans := 0
	for x != 0 {
		if ans < math.MinInt32/10 || ans > math.MaxInt32/10 {
			return 0
		}
		ans = ans*10 + x%10
		x /= 10
	}
	return ans
}

// @lc code=end
