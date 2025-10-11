package leetcode

/*
 * @lc app=leetcode.cn id=1267 lang=golang
 *
 * [1267] 统计参与通信的服务器
 */

// @lc code=start
func countServers(grid [][]int) int {
	row := make([]int, len(grid))
	column := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				row[i]++
				column[j]++
			}
		}
	}

	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 && (row[i] >= 2 || column[j] >= 2) {
				result++
			}
		}
	}

	return result
}

// @lc code=end
