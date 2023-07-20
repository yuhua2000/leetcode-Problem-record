/*
 * @lc app=leetcode.cn id=2428 lang=golang
 *
 * [2428] 沙漏的最大总和
 */

// @lc code=start
func maxSum(grid [][]int) int {
	res := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if ans := grid[i][j] + grid[i-1][j-1] + grid[i-1][j] + grid[i-1][j+1] + grid[i+1][j-1] + grid[i+1][j] + grid[i+1][j+1]; ans > res {
				res = ans
			}
		}
	}
	return res
}

// @lc code=end
