/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	xorSum := 0
	for i := 0; i < len(nums); i++ {
		xorSum ^= nums[i]
	}
	lsp := xorSum ^ -xorSum
	num1, num2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i]^lsp > 0 {
			num1 ^= nums[i]
		} else {
			num2 ^= nums[i]
		}
	}
	return []int{num1, num2}
}

// @lc code=end
