package leetcode

/*
 * @lc app=leetcode.cn id=1260 lang=golang
 *
 * [1260] 二维网格迁移
 */

// @lc code=start
func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for r, row := range grid {
		for j, v := range row {
			index := (r*n + j + k) % (m * n)
			ans[index/n][index%n] = v
		}
	}
	return ans
}

// @lc code=end
