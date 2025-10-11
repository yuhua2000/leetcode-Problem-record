package leetcode

/*
 * @lc app=leetcode.cn id=1124 lang=golang
 *
 * [1124] 表现良好的最长时间段
 */

// @lc code=start
func longestWPI(hours []int) int {
	var n = len(hours)
	var dic = map[int]int{}
	var s = 0
	var res = 0
	for i := 0; i < n; i++ {
		if hours[i] > 8 {
			s += 1
		} else {
			s -= 1
		}
		if s > 0 {
			if res < i+1 {
				res = i + 1
			}
		} else {
			if _, ok := dic[s-1]; ok {
				if res < i-dic[s-1] {
					res = i - dic[s-1]
				}
			}
		}
		if _, ok := dic[s]; !ok {
			dic[s] = i
		}
	}
	return res
}

// @lc code=end
