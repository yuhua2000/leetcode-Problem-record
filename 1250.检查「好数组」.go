/*
 * @lc app=leetcode.cn id=1250 lang=golang
 *
 * [1250] 检查「好数组」
 */

// @lc code=start
func isGoodArray(nums []int) bool {
	var doc func(x, y int) int
	doc = func(x, y int) int {
		if y != 0 {
			return doc(y, x%y)
		}
		return x
	}
	x := nums[0]
	for i := 1; i < len(nums); i++ {
		x = doc(x, nums[i])
		if x == 1 {
			return true
		}
	}
	return x==1
}
// @lc code=end

