package leetcode

/*
 * @lc app=leetcode.cn id=421 lang=golang
 *
 * [421] 数组中两个数的最大异或值
 */

// @lc code=start
func findMaximumXOR(nums []int) (x int) {
	for k := 30; k >= 0; k-- {
		seen := make(map[int]bool)
		for _, num := range nums {
			seen[num>>k] = true
		}
		xNext := x*2 + 1
		found := false

		for _, num := range nums {
			if seen[num>>k^xNext] {
				found = true
				break
			}
		}
		if found {
			x = xNext
		} else {
			x = xNext - 1
		}
	}
	return x
}

// @lc code=end
