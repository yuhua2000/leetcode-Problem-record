/*
 * @lc app=leetcode.cn id=565 lang=golang
 *
 * [565] 数组嵌套
 */

// @lc code=start
func arrayNesting(nums []int) (ans int) {
	n := len(nums)
	for i := range nums {
		cnt := 0
		for nums[i] < n {
			i, nums[i] = nums[i], n //有向图必然由一个或多个环组成,如果在环里存在过 则直接标记即可，也就是说这个环不会产生新的最大值
			cnt++
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return

}

// @lc code=end
