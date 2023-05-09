/*
 * @lc app=leetcode.cn id=2432 lang=golang
 *
 * [2432] 处理用时最长的那个任务的员工
 */

// @lc code=start
func hardestWorker(n int, logs [][]int) int {
	ans, maxCost := logs[0][0], logs[0][1]
	lastCost := maxCost
	for _, log := range logs[1:] {
		idx := log[0]
		cost := log[1] - lastCost
		lastCost = log[1]
		if cost > maxCost || (cost == maxCost && idx < ans) {
			ans = idx
			maxCost = cost
		}
	}
	return ans
}

// @lc code=end
