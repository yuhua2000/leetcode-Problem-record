package leetcode

/*
 * @lc app=leetcode.cn id=944 lang=golang
 *
 * [944] 删列造序
 */

// @lc code=start
func minDeletionSize(strs []string) int {
	result := 0

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				result++
				break
			}
		}
	}
	return result
}

// @lc code=end
