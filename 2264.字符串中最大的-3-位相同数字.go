/*
 * @lc app=leetcode.cn id=2264 lang=golang
 *
 * [2264] 字符串中最大的 3 位相同数字
 */

// @lc code=start
func largestGoodInteger(num string) string {
	result := ""
	for i := 2; i < len(num); i++ {
		if num[i] == num[i-1] && num[i-1] == num[i-2] && num[i-2:i+1] > result {
			result = num[i-2 : i+1]
		}
	}

	return result
}

// @lc code=end
