/*
 * @lc app=leetcode.cn id=2611 lang=golang
 *
 * [2611] 老鼠和奶酪
 */

// @lc code=start
func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	var diff = make([]int, len(reward1))
	result := 0
	for i, v := range reward2 {
		diff[i] = reward1[i] - reward2[i]
		result += v
	}
	sort.Ints(diff)
	n := len(diff)
	for i := 0; i < k; i++ {
		result += diff[n-i-1]
	}
	return result
}

// @lc code=end