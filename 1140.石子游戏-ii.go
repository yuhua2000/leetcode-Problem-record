/*
 * @lc app=leetcode.cn id=1140 lang=golang
 *
 * [1140] 石子游戏 II
 */

// @lc code=start
func stoneGameII(piles []int) int {
	len, sum := len(piles), 0
	dp := make([][]int, len)
	for i := 0; i < len; i++ {
		dp[i] = make([]int, len+1)
	}
	for i := len - 1; i >= 0; i-- {
		sum += piles[i]
		for M := 1; M <= len; M++ {
			if i+2*M >= len {
				dp[i][M] = sum
			} else {
				for x := 1; x <= 2*M; x++ {
					dp[i][M] = max(dp[i][M], sum-dp[i+x][max(M, x)])
				}
			}
		}
	}
	return dp[0][1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
