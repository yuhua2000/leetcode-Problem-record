package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=2500 lang=golang
 *
 * [2500] 删除每行中的最大值
 */

// @lc code=start
func deleteGreatestValue(grid [][]int) (result int) {
	ans := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		sort.Ints(grid[i])
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > ans[j] {
				result += grid[i][j] - ans[j]
				ans[j] = grid[i][j]
			}
		}
	}

	return
}

// @lc code=end
