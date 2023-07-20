/*
 * @lc app=leetcode.cn id=918 lang=golang
 *
 * [918] 环形子数组的最大和
 */

// @lc code=start
func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	leftMax := make([]int, n)
	leftMax[0] = nums[0]
	leftSum, pre, res := nums[0], nums[0], nums[0]
	for i := 1; i < n; i++ {
		pre = max(nums[i], pre+nums[i])
		res = max(res, pre)
		leftSum += nums[i]
		leftMax[i] = max(leftMax[i-1], leftSum)
	}

	rightSum := 0
	for i := n - 1; i > 0; i-- {
		rightSum += nums[i]
		res = max(res, rightSum+leftMax[i-1])
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
