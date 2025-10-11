package leetcode

/*
 * @lc app=leetcode.cn id=2178 lang=golang
 *
 * [2178] 拆分成最多数目的正偶数之和
 */

// @lc code=start
func maximumEvenSplit(finalSum int64) (result []int64) {
	if finalSum%2 != 0 {
		return
	}
	var i int64
	for i = 2; finalSum-i >= 0; i += 2 {
		result = append(result, i)
		finalSum -= i
	}
	result[len(result)-1] += finalSum
	return result
}

// @lc code=end
