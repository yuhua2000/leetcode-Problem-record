package leetcode

func maxValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	ans := make([]int, n)
	ans[0] = grid[0][0]
	for j := 1; j < n; j++ {
		ans[j] = grid[0][j] + ans[j-1]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				ans[j] = max(ans[j-1], ans[j]) + grid[i][j]
			} else {
				ans[j] = ans[j] + grid[i][j]
			}
		}
	}
	return ans[n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
