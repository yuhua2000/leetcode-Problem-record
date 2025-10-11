package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1626 lang=golang
 *
 * [1626] 无矛盾的最佳球队
 */

// @lc code=start
func bestTeamScore1(scores []int, ages []int) int {
	n := len(scores)
	people := make([][2]int, n)
	for i := 0; i < n; i++ {
		people[i] = [2]int{scores[i], ages[i]}
	}
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] < people[j][0] {
			return true
		} else if people[i][0] > people[j][0] {
			return false
		}
		return people[i][1] < people[j][1]
	})
	dp := make([]int, n) // 球队中最大球员序号为 i 的球队最大得分
	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if people[j][1] <= people[i][1] { //前面的球员年纪如果比本球员小或者一样 则可以是这个球员为最大序号 是球队的最大得分 加上 当前球员的分手
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] += people[i][0]
		result = max(result, dp[i])
	}
	return result
}

func bestTeamScore(scores []int, ages []int) int {
	n := len(scores)
	maxArg := 0
	people := make([][2]int, n)
	for i := 0; i < n; i++ {
		people[i] = [2]int{scores[i], ages[i]}
		maxArg = max(maxArg, ages[i])
	}
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] < people[j][0] {
			return true
		} else if people[i][0] > people[j][0] {
			return false
		}
		return people[i][1] < people[j][1]
	})

	t := make([]int, maxArg+1)
	var query func(arg int) int
	query = func(arg int) int {
		ret := 0
		for i := arg; i > 0; i -= lowbit(i) {
			ret = max(ret, t[i])
		}
		return ret
	}
	var update func(arg, cur int)
	update = func(arg, cur int) {
		for i := arg; i <= maxArg; i += lowbit(i) {
			t[i] = max(t[i], cur)
		}
	}
	result := 0
	for i := 0; i < n; i++ {
		score, arg := people[i][0], people[i][1]
		cur := score + query(arg)
		update(arg, cur)
		result = max(cur, result)
	}
	return result
}

func lowbit(x int) int {
	return x & (-x)
}

// @lc code=end
