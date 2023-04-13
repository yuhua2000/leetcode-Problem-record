/*
 * @lc app=leetcode.cn id=2404 lang=golang
 *
 * [2404] 出现最频繁的偶数元素
 */

// @lc code=start
func mostFrequentEven(nums []int) int {
	count := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			count[nums[i]]++
		}
	}
	res, ct := -1, 0
	for k, v := range count {
		if res == -1 || v > ct || (v == ct && k < res) {
			res = k
			ct = v
		}
	}
	return res
}

// @lc code=end
