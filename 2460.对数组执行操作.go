/*
 * @lc app=leetcode.cn id=2460 lang=golang
 *
 * [2460] 对数组执行操作
 */

// @lc code=start
func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}
	index := 0
	for i, num := range nums {
		if num != 0 {
			nums[index] = num
			if index != i {
				nums[i] = 0
			}
			index++
		}
	}
	return nums
}

// @lc code=end
