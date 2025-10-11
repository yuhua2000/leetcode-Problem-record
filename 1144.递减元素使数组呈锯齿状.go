package leetcode

/*
 * @lc app=leetcode.cn id=1144 lang=golang
 *
 * [1144] 递减元素使数组呈锯齿状
 */

// @lc code=start
func movesToMakeZigzag(nums []int) int {
	help := func(ops int) int {
		res := 0
		for i := ops; i < len(nums); i += 2 {
			var a = 0
			if i-1 >= 0 {
				a = max(a, nums[i]-nums[i-1]+1)
			}
			if i+1 < len(nums) {
				a = max(a, nums[i]-nums[i+1]+1)
			}
			res += a
		}
		return res
	}
	return min(help(0), help(1))
}

// @lc code=end
