package main

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	ans := 0
	res := x
	for x != 0 && x > 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	return ans == res
}

// @lc code=end
