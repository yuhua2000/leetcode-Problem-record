/*
 * @lc app=leetcode.cn id=849 lang=golang
 *
 * [849] 到最近的人的最大距离
 */

// @lc code=start
func maxDistToClosest(seats []int) int {
	res := 0
	l := 0
	for l < len(seats) && seats[l] == 0 {
		l++
	}
	res = max(res, l)
	for l < len(seats) {
		r := l + 1
		for r < len(seats) && seats[r] == 0 {
			r++
		}
		if r == len(seats) {
			res = max(r-l-1, res)
		} else {
			res = max(res, (r-l)/2)
		}
		l = r
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
