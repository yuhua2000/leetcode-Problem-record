package leetcode

/*
 * @lc app=leetcode.cn id=1105 lang=golang
 *
 * [1105] 填充书架
 */

// @lc code=start
func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 1000000
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		maxHeight, curWidth := 0, 0
		for j := i; j >= 0; j-- {
			curWidth += books[j][0]
			if curWidth > shelfWidth {
				break
			}
			maxHeight = max(maxHeight, books[j][1])
			dp[i+1] = min(dp[i+1], dp[j]+maxHeight)
		}
	}
	return dp[n]
}

// @lc code=end
