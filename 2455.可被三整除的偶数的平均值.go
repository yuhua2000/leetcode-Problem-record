/*
 * @lc app=leetcode.cn id=2455 lang=golang
 *
 * [2455] 可被三整除的偶数的平均值
 */

// @lc code=start
func averageValue(nums []int) int {
	count := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%6 == 0 {
			count++
			sum += nums[i]
		}
	}
	if count == 0 {
		return 0
	}
	return sum / count
}

// @lc code=end
