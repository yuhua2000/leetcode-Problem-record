package leetcode

/*
 * @lc app=leetcode.cn id=1163 lang=golang
 *
 * [1163] 按字典序排在最后的子串
 */

// @lc code=start
func lastSubstring(s string) string {
	i, j, n := 0, 1, len(s)
	for j < n {
		k := 0
		for j+k < n && s[i+k] == s[j+k] {
			k++
		}
		if j+k < n && s[i+k] < s[j+k] {
			i, j = j, max(j+1, i+k+1)
		} else {
			j = j + k + 1
		}
	}
	return s[i:]
}

// @lc code=end
