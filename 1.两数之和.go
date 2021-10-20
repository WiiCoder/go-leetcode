package main

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
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
