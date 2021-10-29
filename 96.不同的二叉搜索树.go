package main

/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 * 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
 *
 * 示例1:
 * 输入：n = 3
 * 输出：5
 *
 * 示例2:
 * 输入：n = 1
 * 输出：1
 *
 * result：
 * Accepted
 * 19/19 cases passed (0 ms)
 * Your runtime beats 100 % of golang submissions
 * Your memory usage beats 54.71 % of golang submissions (1.9 MB)
 */

// @lc code=start
// 方法二：动态规划
func numTrees2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}

// 方法一：递归，穷举
func numTrees(n int) int {
	memo := make([]int, n+1)
	return helper(n, &memo)
}

func helper(n int, memo *[]int) int {
	if n == 1 || n == 0 {
		return 1
	}
	if (*memo)[n] > 0 {
		return (*memo)[n]
	}
	count := 0
	for i := 0; i <= n-1; i++ {
		count += helper(i, memo) * helper(n-i-1, memo)
	}
	(*memo)[n] = count
	return count
}

// @lc code=end
