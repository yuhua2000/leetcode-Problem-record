/*
 * @lc app=leetcode.cn id=1375 lang=golang
 *
 * [1375] 二进制字符串前缀一致的次数
 */

// @lc code=start
func numTimesAllBlue(flips []int) int {
	n, ans, right := len(flips), 0, 0
	for i := 0; i < n; i++ {
		if right < flips[i] {
			right = flips[i]
		}
		if right == i+1 {
			ans++
		}
	}
	return ans
}

// @lc code=end
