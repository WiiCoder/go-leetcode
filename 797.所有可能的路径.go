package main

/*
 * @lc app=leetcode.cn id=797 lang=golang
 *
 * [797] 所有可能的路径
 */

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	var path = []int{}
	var res = [][]int{}
	var dfs func(s int, path []int)
	dfs = func(s int, path []int) {
		path = append(path, s)
		if s == (len(graph) - 1) {
			var temp = make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			path = []int{}
			return
		}
		for _, v := range graph[s] {
			dfs(v, path)
		}
		path = path[:len(path)-1]
	}
	dfs(0, path)
	return res
}

// @lc code=end
