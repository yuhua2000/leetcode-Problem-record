package leetcode

/*
 * @lc app=leetcode.cn id=1911 lang=golang
 *
 * [1911] 最大子序列交替和
 */

// @lc code=start
func maxAlternatingSum(nums []int) int64 {
	even, odd := nums[0], 0
	for i := 0; i < len(nums); i++ {
		even = max(even, odd+nums[i])
		odd = max(odd, even-nums[i])
	}
	return int64(even)

}

// @lc code=end
