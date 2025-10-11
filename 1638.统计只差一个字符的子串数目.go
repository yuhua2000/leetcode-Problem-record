package leetcode

/*
 * @lc app=leetcode.cn id=1638 lang=golang
 *
 * [1638] 统计只差一个字符的子串数目
 */

// @lc code=start
func countSubstrings(s string, t string) int {
	m := len(s)
	n := len(t)
	var dpl = make([][]int, m+1)
	var dpr = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dpl[i] = make([]int, n+1)
		dpr[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i] == t[j] {
				dpl[i+1][j+1] = dpl[i][j] + 1
			} else {
				dpl[i+1][j+1] = 0
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dpr[i][j] = dpl[i+1][j+1] + 1
			} else {
				dpr[i][j] = 0
			}
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i] != t[j] {
				ans += (dpl[i][j] + 1) * (dpr[i+1][j+1] + 1)
			}
		}
	}
	return ans
}

// @lc code=end
