/*
 * @lc app=leetcode.cn id=1615 lang=golang
 *
 * [1615] 最大网络秩
 */

// @lc code=start
func maximalNetworkRank(n int, roads [][]int) int {
	ans := make([]int, n)
	result := 0
	unicom := make(map[[2]int]bool)
	for _, road := range roads {
		x := road[0]
		y := road[1]
		ans[x]++
		ans[y]++
		unicom[[2]int{x, y}] = true
	}
	for i := 0; i < len(ans); i++ {
		for j := i + 1; j < len(ans); j++ {
			if unicom[[2]int{i, j}] || unicom[[2]int{j, i}] {
				if ans[i]+ans[j]-1 > result {
					result = ans[i] + ans[j] - 1
				}
			} else {
				if ans[i]+ans[j] > result {
					result = ans[i] + ans[j]
				}
			}
		}
	}
	return result
}

// @lc code=end
