/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	ans := [][]int{}
	l, r := newInterval[0], newInterval[1]
	merge := false
	for _, v := range intervals {
		if v[0] > r {
			if !merge {
				ans = append(ans, []int{l, r})
				merge = true
			}
			ans = append(ans, v)
		} else if v[1] < l {
			ans = append(ans, v)
		} else {
			l = min(l, v[0])
			r = max(r, v[1])
		}
	}
	if !merge {
		ans = append(ans, []int{l, r})
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
