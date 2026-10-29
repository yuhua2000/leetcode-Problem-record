package leetcode

func minCost(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([][1024]bool, m)
	var tempZeroArr [1024]bool

	dp[0][grid[0][0]] = true
	var next [1024]bool

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}

			clear(next[:])
			up, left := tempZeroArr, tempZeroArr
			if i > 0 {
				up = dp[j]
			}
			if j > 0 {
				left = dp[j-1]
			}

			for k := 0; k < 1024; k++ {
				if up[k] || left[k] {
					next[grid[i][j]^k] = true
				}
			}

			dp[j] = next
		}
	}

	for xor := 0; xor < 1024; xor++ {
		if dp[m-1][xor] {
			return xor
		}
	}

	return -1
}
