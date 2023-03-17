/*
 * @lc app=leetcode.cn id=2389 lang=golang
 *
 * [2389] 和有限的最长子序列
 */

// @lc code=start
func answerQueries(nums []int, queries []int) []int {
	n := len(nums)
	m := len(queries)
	sort.Ints(nums)
	count := 0
	for i := 0; i < n; i++ {
		count += nums[i]
		nums[i] = count
	}
	ans := make([]int, m)
	for i := 0; i < m; i++ {
		ans[i] = search(nums, queries[i])
	}
	return ans
}

func search(m []int, x int) int {
	l, r := 0, len(m)
	for l < r {
		mid := (l + r) >> 1
		if m[mid] <= x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// @lc code=end
