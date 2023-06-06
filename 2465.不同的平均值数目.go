/*
 * @lc app=leetcode.cn id=2465 lang=golang
 *
 * [2465] 不同的平均值数目
 */

// @lc code=start
func distinctAverages(nums []int) int {
	res := make(map[int]struct{})
	sort.Ints(nums)
	for len(nums) > 0 {
		res[nums[0]+nums[len(nums)-1]] = struct{}{}
		nums = nums[1 : len(nums)-1]
	}
	return len(res)
}

// @lc code=end
