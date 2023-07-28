/*
 * @lc app=leetcode.cn id=2050 lang=golang
 *
 * [2050] 并行课程 III
 */

// @lc code=start
func minimumTime(n int, relations [][]int, time []int) (result int) {
	prev := make([][]int, n+1)
	for i := 0; i < len(relations); i++ {
		prev[relations[i][0]] = append(prev[relations[i][0]], relations[i][1])
	}
	memo := make(map[int]int, 0)
	var dp func(int) int
	dp = func(i int) int {
		if cur, ok := memo[i]; !ok {
			for _, p := range prev[i] {
				cur = max(cur, dp(p))
			}
			cur += time[i-1]
			memo[i] = cur
			return cur
		} else {
			return cur
		}
	}
	for i := 1; i <= n; i++ {
		result = max(result, dp(i))
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
