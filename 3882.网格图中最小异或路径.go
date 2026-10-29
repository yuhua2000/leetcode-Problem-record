package leetcode

func minCost(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([][][1024]bool, n)
	var tempZeroArr [1024]bool

	for i := 0; i < n; i++ {
		dp[i] = make([][1024]bool, m)
	}
	dp[0][0][grid[0][0]] = true

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}

			up, left := tempZeroArr, tempZeroArr
			if i > 0 {
				up = dp[i-1][j]
			}
			if j > 0 {
				left = dp[i][j-1]
			}

			for k := 0; k < 1024; k++ {
				if up[k] || left[k] {
					dp[i][j][grid[i][j]^k] = true
				}
			}
		}
	}

	for xor := 0; xor < 1024; xor++ {
		if dp[n-1][m-1][xor] {
			return xor
		}
	}

	return -1
}
