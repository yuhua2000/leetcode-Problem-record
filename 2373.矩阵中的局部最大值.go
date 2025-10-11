package leetcode

/*
 * @lc app=leetcode.cn id=2373 lang=golang
 *
 * [2373] 矩阵中的局部最大值
 */

// @lc code=start
func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	ans := make([][]int, n-2)
	for i := 1; i <= len(ans); i++ {
		row := make([]int, n-2)
		for j := 1; j <= len(row); j++ {
			mx := grid[i][j]
			for r := i - 1; r <= i+1; r++ {
				for c := j - 1; c <= j+1; c++ {
					if grid[r][c] > mx {
						mx = grid[r][c]
					}
				}
			}
			row[j-1] = mx
		}
		ans[i-1] = row
	}
	return ans
}

// @lc code=end
