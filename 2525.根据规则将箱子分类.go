/*
 * @lc app=leetcode.cn id=2525 lang=golang
 *
 * [2525] 根据规则将箱子分类
 */

// @lc code=start
func categorizeBox(length int, width int, height int, mass int) string {
	var heavy, bulky string
	if mass >= 100 {
		heavy = "Heavy"
	}

	if length*width*height >= 1e9 || length >= 1e4 || width >= 1e4 || height >= 1e4 {
		bulky = "Bulky"
	}
	if len(heavy)+len(bulky) == len(heavy) || len(heavy)+len(bulky) == len(bulky) {
		if len(heavy)+len(bulky) == 0 {
			return "Neither"
		}
		return bulky + heavy
	} else {
		return "Both"
	}
}

// @lc code=end
