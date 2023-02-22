/*
 * @lc app=leetcode.cn id=1498 lang=golang
 *
 * [1498] 满足条件的子序列数目
 */

// @lc code=start
func numSubseq(nums []int, target int) int {
	const mod = 1000000007
	sort.Ints(nums)
	tmp := make([]int, len(nums))
	tmp[0] = 1
	for i := 1; i < len(tmp); i++ {
		tmp[i] = (tmp[i-1] << 1) % mod
	}
	ret := 0
	l := 0
	r := len(nums) - 1
	for l <= r {
		if nums[l]+nums[r] > target {
			r--
		} else {
			ret = (ret + tmp[r-l]) % mod
			l++
		}
	}
	return ret
}

// @lc code=end

