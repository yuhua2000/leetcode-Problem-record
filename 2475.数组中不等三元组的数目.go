package leetcode

/*
 * @lc app=leetcode.cn id=2475 lang=golang
 *
 * [2475] 数组中不等三元组的数目
 */

// @lc code=start
func unequalTriplets(nums []int) int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	res, n, t := 0, len(nums), 0
	for _, v := range count {
		res, t = res+v*t*(n-v-t), t+v
	}
	return res
}

// @lc code=end
