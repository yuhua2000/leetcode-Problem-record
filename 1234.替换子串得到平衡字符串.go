/*
 * @lc app=leetcode.cn id=1234 lang=golang
 *
 * [1234] 替换子串得到平衡字符串
 */

// @lc code=start
func balancedString(s string) int {
	cnt := map[byte]int{}
	for i := range s {
		cnt[s[i]]++
	}
	partial := len(s) / 4
	check := func() bool {
		if cnt['Q'] > partial || cnt['W'] > partial || cnt['E'] > partial || cnt['R'] > partial {
			return false
		}
		return true
	}
	if check() {
		return 0
	}
	res := len(s)
	r := 0
	for l := range s {
		for r < len(s) && !check() {
			cnt[s[r]]--
			r++
		}
		if !check() {
			break
		}
		if r-l < res {
			res = r - l
		}
		cnt[s[l]]++
	}
	return res
}
// @lc code=end

