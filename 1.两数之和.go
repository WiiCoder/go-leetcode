package main

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和(Easy)
 *
 * 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
 * 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
 * 你可以按任意顺序返回答案。
 *
 * 示例 1：
 * 输入：nums = [2,7,11,15], target = 9
 * 输出：[0,1]
 * 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
 *
 * result:
 * Accepted
 * 55/55 cases passed (32 ms)
 * Your runtime beats 17.59 % of golang submissions
 * Your memory usage beats 99.67 % of golang submissions (3.6 MB)
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// 双for、暴力破解
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// @lc code=end
