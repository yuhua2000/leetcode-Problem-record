package leetcode

/*
 * @lc app=leetcode.cn id=2278 lang=golang
 *
 * [2278] 字母在字符串中的百分比
 */

// @lc code=start
func percentageLetter(s string, letter byte) int {
	var numerator int
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			numerator++
		}
	}

	return 100 * numerator / len(s)
}

// @lc code=end
