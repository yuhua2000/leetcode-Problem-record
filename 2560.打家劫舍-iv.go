package leetcode

/*
 * @lc app=leetcode.cn id=2560 lang=golang
 *
 * [2560] 打家劫舍 IV
 */

// @lc code=start
func minCapability(nums []int, k int) int {
	lower, upper := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		lower = min(lower, nums[i])
		upper = max(upper, nums[i])
	}

	for lower <= upper {
		middle := (lower + upper) / 2
		count, visited := 0, false
		for _, x := range nums {
			if x <= middle && !visited {
				count, visited = count+1, true
			} else {
				visited = false
			}
		}
		if count >= k {
			upper = middle - 1
		} else {
			lower = middle + 1
		}
	}
	return lower
}

// @lc code=end
