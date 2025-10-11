package leetcode

/*
 * @lc app=leetcode.cn id=2379 lang=golang
 *
 * [2379] 得到 K 个黑块的最少涂色次数
 */

// @lc code=start
func minimumRecolors(blocks string, k int) int {
	w := 0
	ans := 100
	for i, c := range blocks {
		if c == 'W' {
			w++
		}
		if i >= k-1 {
			ans = min(ans, w)
			if blocks[i-k+1] == 'W' {
				w--
			}
		}
	}
	return ans
}

// @lc code=end
