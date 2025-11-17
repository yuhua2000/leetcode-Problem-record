package leetcode

/*
 * @lc app=leetcode.cn id=1437 lang=golang
 *
 * [1437] 是否所有 1 都至少相隔 k 个元素
 */

// @lc code=start
func kLengthApart(nums []int, k int) bool {
	prefOneIndex := -k - 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if i-prefOneIndex-1 < k {
				return false
			}
			prefOneIndex = i
		}
	}
	return true
}

// @lc code=end
