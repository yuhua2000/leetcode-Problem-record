package leetcode

/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte) {
	for l, r := 0, len(s)-1; l < r; l++ {
		s[l], s[r] = s[r], s[l]
		r--
	}
}

// @lc code=end
