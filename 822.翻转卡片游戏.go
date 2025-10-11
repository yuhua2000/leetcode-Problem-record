package leetcode

/*
 * @lc app=leetcode.cn id=822 lang=golang
 *
 * [822] 翻转卡片游戏
 */

// @lc code=start
func flipgame(fronts []int, backs []int) int {
	n := len(fronts)
	same := make(map[int]bool)
	for i := 0; i < n; i++ {
		if fronts[i] == backs[i] {
			same[fronts[i]] = true
		}
	}
	res := 3000
	for _, x := range fronts {
		if x < res && !same[x] {
			res = x
		}
	}
	for _, x := range backs {
		if x < res && !same[x] {
			res = x
		}
	}
	return res % 3000
}

// @lc code=end
