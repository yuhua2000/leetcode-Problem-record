package leetcode

/*
 * @lc app=leetcode.cn id=2270 lang=golang
 *
 * [2270] 分割数组的方案数
 */

// @lc code=start
func waysToSplitArray(nums []int) int {
	var l, r int64
	for _, num := range nums {
		r += int64(num)
	}

	var result int
	for i := 0; i < len(nums)-1; i++ {
		l += int64(nums[i])
		r -= int64(nums[i])
		if l >= r {
			result++
		}
	}
	return result
}

// @lc code=end
