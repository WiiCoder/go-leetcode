package main

/*
 * @lc app=leetcode.cn id=797 lang=golang
 *
 * [797] 所有可能的路径(Medium)
 *
 * 给你一个有 n 个节点的 有向无环图（DAG），请你找出所有从节点 0 到节点 n-1 的路径并输出（不要求按特定顺序）
 * 二维数组的第 i 个数组中的单元都表示有向图中 i 号节点所能到达的下一些节点，空就是没有下一个结点了。
 * 译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a 。
 *
 * 示例 1：
 * 输入：graph = [[1,2],[3],[3],[]]
 * 输出：[[0,1,3],[0,2,3]]
 * 解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3
 *
 * result:
 * Accepted
 * 30/30 cases passed (8 ms)
 * Your runtime beats 83.33 % of golang submissions
 * Your memory usage beats 41.15 % of golang submissions (6.7 MB)
 */

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	// 路径
	var path = []int{}
	// 所有路径
	var res = [][]int{}
	// 深度
	var dfs func(s int, path []int)
	dfs = func(s int, path []int) {
		path = append(path, s)
		if s == (len(graph) - 1) {
			var temp = make([]int, len(path))
			// 拷贝路径
			copy(temp, path)
			res = append(res, temp)
			path = []int{}
			return
		}
		// 递归
		for _, v := range graph[s] {
			dfs(v, path)
		}
		path = path[:len(path)-1]
	}
	dfs(0, path)
	return res
}

// @lc code=end
