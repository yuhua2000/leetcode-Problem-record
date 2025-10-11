package leetcode

/*
 * @lc app=leetcode.cn id=1749 lang=golang
 *
 * [1749] 任意子数组和的绝对值的最大值
 */

// @lc code=start
func maxAbsoluteSum(nums []int) int {
	positiveMax, negativeMin := 0, 0
	positiveSum, negativeSum := 0, 0
	for _, num := range nums {
		positiveSum += num
		positiveMax = max(positiveMax, positiveSum)
		positiveSum = max(0, positiveSum)

		negativeSum += num
		negativeMin = min(negativeMin, negativeSum)
		negativeSum = min(0, negativeSum)
	}
	return max(positiveMax, -negativeMin)
}

// @lc code=end
