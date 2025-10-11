package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=1016 lang=golang
 *
 * [1016] 子串能表示从 1 到 N 数字的二进制串
 */

// @lc code=start
func queryString(s string, n int) bool {
	for i := 1; i <= n; i++ {
		if !strings.Contains(s, strconv.FormatUint(uint64(i), 2)) {
			return false
		}
	}
	return true
}

// @lc code=end
